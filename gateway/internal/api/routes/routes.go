package routes

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(authRoutes AuthRoutes) *Routes {
	return &Routes{
		authRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
