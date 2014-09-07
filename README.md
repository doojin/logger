Golang Logger
=============

[![Build Status](https://travis-ci.org/doojin/logger.svg)](https://travis-ci.org/doojin/logger)

Usage
=====

The simplest way to use logger:

    import "./logger"
    
    func main() {
        logger.Info("Hello, world!")
    }
Will output:
>15:52:57 [INFO] Hello, world!

There are 6 message "levels" supported by logger currently:

* Info
* Warning
* Error
* Fatal
* Debug
* Trace

There are 6 methods available to log messages with corresponding levels:

    func main() {
        logger.Info("Here goes info message")
        logger.Warn("Here goes warning message")
        logger.Error("Here goes error message")
        logger.Fatal("Here goes fatal message")
        logger.Debug("Here goes debug message")
        logger.Trace("Here goes trace message")
    }
    
There are also 6 methods to log messages with the new line symbol (\n) in the end of the message:

    func main() {
        logger.Infoln("Here goes info message")
        logger.Warnln("Here goes warning message")
        logger.Errorln("Here goes error message")
        logger.Fatalln("Here goes fatal message")
        logger.Debugln("Here goes debug message")
        logger.Traceln("Here goes trace message")
    }
    
And 6 methods which allows to format messages before logging them:

    func main() {
        logger.Infof("Here goes %v message", "info")
        logger.Warnf("Here goes %v message", "warning")
        logger.Errorf("Here goes %v message", "error")
        logger.Fatalf("Here goes %v message", "fatal")
        logger.Debugf("Here goes %v message", "debug")
        logger.Tracef("Here goes %v message", "trace")
    }
    
These methods also puts new line symbols in the end of the message.