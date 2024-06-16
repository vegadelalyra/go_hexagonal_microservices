package shortener

type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(rediret *Redirect) error
}