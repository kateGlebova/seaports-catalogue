package mongo

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/kateGlebova/seaports-catalogue/pkg/lifecycle"
)

// Run creates MongoDB session
func (r *Repository) Run() {
	log.Print("Dialing MongoDB...")
	session, err := mgo.Dial(r.url)
	if err != nil {
		r.err = err
		lifecycle.KillTheApp()
	}
	r.session = session
}

// Stop closes MongoDB session if one exists
func (r *Repository) Stop() error {
	if r.err != nil {
		return r.err
	}
	log.Print("Closing MongoDB session...")
	if r.session != nil {
		r.session.Close()
		return nil
	}
	return nil
}
