package database

func (c *Connector) Close() {
	err := c.Orm.Close()
	c.logger.CheckErr(err)
	c.logger.Info("connection to database closed")
}
