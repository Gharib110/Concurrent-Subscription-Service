package handlers

import (
	"database/sql"
	"github.com/Gharib110/Concurrent-Subscription-Service/data"
	"github.com/alexedwards/scs/v2"
	"log"
	"sync"
)

type Config struct {
	Session *scs.SessionManager
	DB      *sql.DB
	InfoLog *log.Logger
	ErrLog  *log.Logger
	Wait    *sync.WaitGroup
	Data    data.Models
}
