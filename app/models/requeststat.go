package models

type RequestStat struct {
	Request FizzBuzzRequest  `json:"request"`
	HitCount int                    `json:"hitCount""`
}