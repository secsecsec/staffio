package ldap

import (
	// "fmt"
	"log"

	"lcgc/platform/staffio/models"
)

func Modify(uid, password string, staff *models.Staff) (err error) {
	for _, ls := range ldapSources {
		err = ls.Modify(uid, password, staff)
		if err != nil {
			log.Printf("Modify at %s ERR: %s", ls.Addr, err)
		}
	}
	return
}

func (ls *LdapSource) Modify(uid, password string, staff *models.Staff) error {
	if ls.Debug {
		log.Printf("change profile for %s staff: %v", uid, staff)
	}
	userdn := ls.UDN(uid)
	err := ls.Bind(userdn, password, true)
	if err != nil {
		return ErrLogin
	}
	entry, err := ls.getEntry(userdn)
	if err != nil {
		return err
	}

	modify := makeModifyRequest(userdn, entry, staff)

	if err := ls.c.Modify(modify); err != nil {
		log.Printf("Modify ERROR: %s\n", err.Error())
	}

	return nil
}
