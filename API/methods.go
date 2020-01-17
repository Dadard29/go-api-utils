package API

func (a API) Start() {
	a.Service.Start()
}

func (a API) Stop() {
	err := a.Service.Stop()
	a.Logger.CheckErr(err)

	a.Database.Close()
}