package urlrepo

type URLRepository interface {
	GetUrlById(id string) (string, error)
	AddUrl(url string) (string, error)
}
