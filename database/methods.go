package database

func (c *Connector) Close() {
	if c != nil {
		err := c.Orm.Close()
		c.logger.CheckErr(err)
		c.logger.Info("connection to database closed")
	}
}
