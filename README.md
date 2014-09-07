Golang Simple Logger
====================

[![Build Status](https://travis-ci.org/doojin/logger.svg)](https://travis-ci.org/doojin/logger)

I am glad to introduce you a very simple implementation of Go Logger.
The main idea of the project is to create a logger which will be really easy to use and configure (if need)

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
For more information about message formatting please read information provided in documentation for "fmt" package: http://golang.org/pkg/fmt/

There is also logger.Log method available, which accept 2 arguments: Level id, message and arguments for message formatting:

    func main() {
        logger.Log("info", "Hello, %v!", "world")
        logger.Log("info", "Hello, world!")
    }
    
This will output twice the same message:

>16:40:53 [INFO] Hello, world!
>
>16:40:53 [INFO] Hello, world!

as level id you can pass "info", "warn", "error", "fatal", "debug", "trace".
Note, that Log() method adds new line symbol in the end of the message.

Before logging information you may want to configure the logger.
For example, you can configure the default output format (default format is: "{time} [{level}] {message}").
To do this, you can change value of logger.Settings.Layout.

    func main() {
        logger.Settings.Layout = "{message} *** {time} *** {level}"
        logger.Infoln("Test message 1")
        
        logger.Settings.Layout = "{message}"
        logger.Infoln("Test message 2")
    }
    
The output will be:
>Test message 1 *** 16:17:29 *** INFO
>
>Test message 2

There property logger.Settings.TimeFormat is responsible for time format in your outputs.
By default, format is: hh:mm:ss, where hh are hours, mm - minutes and ss - seconds.

    func main() {
        logger.Settings.TimeFormat = "Mon, 02 Jan 2006 15:04:05"
        logger.Infoln("Custom time format")
    }

Will output:
>Sun, 07 Sep 2014 16:25:17 [INFO] Custom time format

For more information about time formatting please read documentation: http://golang.org/src/pkg/time/format.go

Finally, if you want to output log message to the file instead of console, you should change logger.Settings.Writer and logger.Settings.Filename properties:

    func main() {
        logger.Settings.Writer = "file"
        logger.Settings.Filename = "log_file.txt"
        logger.Infoln("Output")
    }
    
It will create log_file.txt file and will add to the end:

>16:30:21 [INFO] Output

If file does not exist, logger will try to create it.
By default, logger.Settings.Writer can accept "default", "console" and "file" values. Default writer is Console writer. So setting "default" and "console" to the logger.Settings.Writer is the same thing.

Sometimes you will need to create multiple loggers with different settings. To create new logger, you can use logger.New() method, which returns pointer to the new created logger (*Logger):

    func main() {
        myConsoleLogger := logger.New()
        myConsoleLogger.Settings.Layout = "{time} {message}"
        myConsoleLogger.Infoln("Hello, world!")
        
        myFileLogger := logger.New()
        myFileLogger.Settings.Writer = "file"
        myFileLogger.Settings.Filename = "logs.txt"
        myFileLogger.Settings.TimeFormat = "15:04"
        myFileLogger.Infoln("Hello, world!")
    }
    
myConsoleLogger and myFileLogger are two different loggers with custom settings. The first one outputs to console, the second one to file.