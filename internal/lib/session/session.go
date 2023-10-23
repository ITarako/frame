package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

const cookieName = "apisessid"
const keyUserID = "user"

type Session struct {
	store sessions.Store
}

func NewSession() *Session {
	key := []byte("secret-key") //todo config
	store := sessions.NewCookieStore(key)

	return &Session{store: store}
}

func (s *Session) SetData(w http.ResponseWriter, r *http.Request, key string, data any) error {
	sess, err := s.store.Get(r, cookieName)
	if err != nil {
		return err
	}

	sess.Values[key] = data
	return sess.Save(r, w)
}

func (s *Session) SetUserID(w http.ResponseWriter, r *http.Request, userID int32) error {
	return s.SetData(w, r, keyUserID, userID)
}

func (s *Session) GetUserID(r *http.Request) (int32, error) {
	sess, err := s.store.Get(r, cookieName)
	if err != nil {
		return 0, err
	}

	val := sess.Values[keyUserID]
	userID, ok := val.(int32)
	if !ok {
		return 0, nil
	}

	return userID, nil
}

func (s *Session) Logout(w http.ResponseWriter, r *http.Request) error {
	sess, err := s.store.Get(r, cookieName)
	if err != nil {
		return err
	}

	sess.Values[keyUserID] = 0
	sess.Options.MaxAge = -1
	return sess.Save(r, w)
}
