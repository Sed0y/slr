package models

import (
	"archive/zip"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Конвертирует дату в формат Ms Access
// формат даты в аксесе - #ММ/ДД/ГГГГ#
//
// возвращает строку
//
func ConvertTimeToAccessDate(datetime time.Time) string {

	var result string
	result = "#" +
		strconv.Itoa(int(datetime.Month())) + "/" +
		strconv.Itoa(datetime.Day()) + "/" +
		strconv.Itoa(datetime.Year()) + "#"

	return result
}

// Преобразует строку для запроса
// если строка пустая, то возвращает - NULL
// если есть значение, то - 'value'
//
func Query_StringOrNull(value string) string {

	value = strings.Replace(value, "'", " ", -1)
	query := ""
	if value == "" {
		query += "NULL "
	} else {
		query += "'" + value + "' "
	}
	return query
}

// Преобразует строку для запроса
// если строка пустая, то возвращает - NULL
// если есть значение, то конвертирует в формат даты PostgreSQL
// пример: 21.03.2018 -> 2018-03-21
//
func Query_StringOrNull_ToPostgresDate(value string) string {

	value = strings.Replace(value, "'", " ", -1)
	query := ""
	if value == "" {
		query += "NULL "
	} else {
		split_date := strings.Split(value, ".")
		query += "'" + split_date[2] + "-" + split_date[1] + "-" + split_date[0] + "'"
	}
	return query

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Строка из 8-ми случайных символов
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Архивация файлов
func ZipFiles(filename string, files []string) error {

	newfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, zipfile)
		if err != nil {
			return err
		}
	}
	return nil
}
