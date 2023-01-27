package handlers

import (
	"database/sql"
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
}
