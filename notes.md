receive ctx + logger fields, return child context cancel func (for use as a flush) and a request ID (uuid - should really colide tbh)

internally => map[requestID][logger / fields]

logger.Error(string)
logger.Info(string)
logger.Debug(string)

