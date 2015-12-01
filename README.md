#GOST (GO STandards) Logger #

[![Build Status](https://travis-ci.org/Barberrrry/gost.svg?branch=master)](https://travis-ci.org/Barberrrry/gost)

```import "github.com/go-st/logger"```

This library contains common logger interface and simple implementations for it. Use interface ILogger in your library as input parameter, if you would like to log some messages.

### logger.ILogger ###

This interface uses levels determined by RFC 5424 (https://tools.ietf.org/html/rfc5424). Use interface logger.ILogger in your library as input parameter, if you would like to log some messages. It will be easy to change logger for any other supporting this interface.

### logger.BuiltinLogger ###

logger.BuiltinLogger is an adapter for built-in package "log". You can just create instance of log.Logger and pass it to BuiltinLogger. BuiltinLogger filter messages by level, but output them simply by log.Logger.

### logger.NilLogger ###

logger.NilLogger does nothing. Use it for your tests or if you just don't want to see any log messages.
