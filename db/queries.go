package db

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/noilpa/rest/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// TODO: protect from SQL injection

var ErrOffset = errors.New("offset greater than amount film's ")

// Construct query for database and execute it
// Implement filters and pagination
func Films(size, offset uint, date, genre string) ([]utils.Film, error) {

	// SELECT * FROM films
	// WHERE films.genre IN (genre[0], genre[1] ...)
	//   AND films.date date[0] date[1]

	conditions, err := prepareFilmsConditions(strings.Fields(date), strings.Fields(genre))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM films %s", conditions)
	rows, err := DB.Query(query)
	if err != nil {
		// err no db connection
		return nil, err
	}

	var filmsList []utils.Film
	for rows.Next() {
		var f utils.Film
		err := rows.Scan(&f.Id, &f.Name, &f.Date, &f.Genre)

		if err != nil {
			// if we scan corrupted row
			fmt.Println(err)
			continue
		}
		filmsList = append(filmsList, f)
	}


	// pagination
	last := offset + size
	length := uint(len(filmsList))
	if offset > length {
		return nil, ErrOffset
	}

	if last > length {
		last = length
	}
	filmsList = filmsList[offset: last]

	// final result


	return filmsList, nil

}

// Return only valid genres from url parameters
func validateGenres(genre []string) ([]string, error) {

	query := "SELECT name FROM genres"
	rows, err := DB.Query(query)
	if err != nil {
		// err no db connection
		return nil, err
	}

	var genresList []string
	for rows.Next() {
		var r string
		err := rows.Scan(&r)
		if err != nil {
			// if we scan corrupted row log it
			fmt.Println(err)
			continue
		}
		genresList = append(genresList, r)
	}

	// get only valid genres from url
	var validGenres []string
	for _, g := range genre {
		if utils.Contains(genresList, g) {
			// add quotes to genre for sql query
			validGenres = append(validGenres, fmt.Sprintf("'%s'", g))
		}
	}
	return validGenres, nil
}

// Check sign for valid value =, >, <, >=, <=
func validateSign(sign string) bool {
	validSigns := []string{"=", ">", "<", ">=", "<=",}
	if utils.Contains(validSigns, sign) {
		return true
	}
	return false
}

// Check date for valid value YYYY-MM-DD
func validateDate(date string) bool {

	tokens := strings.Split(date, "-")
	if len(tokens) == 3 {
		return true
	}

	fmt.Println("Expected YYYY-MM-DD, received", date)
	return false

}

// Build condition string for query
func prepareFilmsConditions(date , genre []string) (string, error) {

	var dateCondition string
	if len(date) != 0 {

		// if request consist of only date
		// unification
		if len(date) == 1 {
			date = append(date, date[0])
		}

		if !validateSign(date[0]) {
			date[0] = "="
		}

		if validateDate(date[1]) {
			dateCondition = fmt.Sprintf("date %s '%s'", date[0], date[1])
		}
	}

	var genreCondition string
	genre, err := validateGenres(genre)
	if err != nil {
		return "", err
	}
	if len(genre) != 0 {
		genreCondition = fmt.Sprintf("genre IN (%s)", fmt.Sprint(strings.Join(genre, ", ")))
	}

	var condition string
	if genreCondition != "" && dateCondition != "" {
		condition = fmt.Sprintf("WHERE %s AND %s", genreCondition, dateCondition)
	} else if genreCondition != "" {
		condition = fmt.Sprintf("WHERE %s", genreCondition)
	} else if dateCondition != "" {
		condition = fmt.Sprintf("WHERE %s", dateCondition)
	}

	return condition, nil

}

// Construct query for add user in database
// return user_id in success case else err message
func Registrations(user utils.UserInfo) (map[string]interface{}, error) {

	passwordHash, err := hashPassword(user.Usr.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = DB.QueryRow("INSERT INTO users(login, password) VALUES ($1, $2) RETURNING id",
		user.Usr.Login, passwordHash).Scan(&user.Usr.Id)

	//if err != nil {
	//	fmt.Println(err)
	//	return "", err
	//}

	// get id for next insert and returning

	processUserInfoValues(&user)

	_, err = DB.Query("INSERT INTO users_info(user_id, name, age, phone) VALUES ($1, $2, $3, $4)",
		user.Usr.Id, user.Name, user.Age, user.Phone)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	id := map[string]interface{} {"user_id": user.Usr.Id}

	return id, nil
}

func processUserInfoValues(u *utils.UserInfo) {

	// if username is not specified, use login
	if u.Name == "" {
		u.Name = u.Usr.Login
	}

	// change +7 to 8 for unification (+7-123-456 -> 8-123-456)
	if strings.HasPrefix(u.Phone, "+7") {
		u.Phone = fmt.Sprintf("8%s", u.Phone[2:])
	}

	// remove all "-" chars (8-123-456 -> 8123456)
	u.Phone = strings.Join(strings.Split(u.Phone,"-"), "")

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// token is string or []byte
func Authorizations(user utils.User) (bool, error) {

	// SELECT password FROM users WHERE login = $1
	var u utils.User
	err := DB.QueryRow("SELECT password FROM users WHERE login = $1", user.Login).Scan(&u.Password)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return checkPasswordHash(user.Password, u.Password), nil

}

//var insertUser = "INSERT INTO User(login, password) VALUES ($1, $2)"
//
//// SELECT * FROM User
//func Insert(row utils.User) uint {
//
//	var id uint
//	DB.Exec(insertUser, row.Login, row.Password)
//
//
//	return id
//}