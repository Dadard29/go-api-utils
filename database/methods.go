package database

func (c *Connector) Close() {
	// fixme ?
	if c != nil {
		c.logger.Info("connection to database closed")
	}
}
