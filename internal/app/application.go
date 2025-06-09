package app

type Service struct {
}

type Store struct {
}

type Application struct {
	service *Service
	store   *Store
}

func NewApplication(service *Service, store *Store) *Application {
	return &Application{
		service: service,
		store:   store,
	}
}

func (app *Application) Service() *Service {
	return app.service
}

func (app *Application) Store() *Store {
	return app.store
}
