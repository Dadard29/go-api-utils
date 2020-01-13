package API

func (a API) Start() {
	a.Service.Start()
}

func (a API) Stop() {
	a.Service.Stop()
	a.Connector.Close()
}