package shortener

type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(rediret *Redirect) error
}

