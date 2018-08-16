package cli

import (
    "fmt"
    "os"
)

type Color string

const (
    colorPrefix = "\x1b["

    DefaultColor Color = colorPrefix + "39m"
    Black        Color = colorPrefix + "30m"
    Red          Color = colorPrefix + "31m"
    Green        Color = colorPrefix + "32m"
    Yellow       Color = colorPrefix + "33m"
    Blue         Color = colorPrefix + "34m"
    Magenta      Color = colorPrefix + "35m"
    Cyan         Color = colorPrefix + "36m"
    LightGray    Color = colorPrefix + "37m"
    DarkGray     Color = colorPrefix + "90m"
    LightRed     Color = colorPrefix + "91m"
    LightGreen   Color = colorPrefix + "92m"
    LightYellow  Color = colorPrefix + "93m"
    LightBlue    Color = colorPrefix + "94m"
    LightMagenta Color = colorPrefix + "95m"
    LightCyan    Color = colorPrefix + "96m"
    White        Color = colorPrefix + "97m"
)

var errorColor = LightRed
var warnColor = Yellow
var infoColor = White
var info2Color = LightGray
var successColor = LightGreen
var debugColor = LightBlue
var traceColor = LightMagenta
var importantColor = LightYellow

func SetError(color Color) {
    errorColor = color
}
func SetWarn(color Color) {
    warnColor = color
}
func SetInfo(color Color) {
    infoColor = color
}
func SetSuccess(color Color) {
    successColor = color
}
func SetDebug(color Color) {
    debugColor = color
}
func SetImportant(color Color) {
    importantColor = color
}
func SetInfo2(color Color) {
    info2Color = color
}
func SetTrace(color Color) {
    traceColor = color
}

func SetColor(file *os.File, color Color) {
    fmt.Fprint(file, color)
}

func UnsetColor(file *os.File) {
    fmt.Fprint(file, DefaultColor)
}

func Fprintf(file *os.File, color Color, format string, args ...interface{}) (n int, err error) {
    SetColor(file, color)
    defer UnsetColor(file)
    return fmt.Fprintf(file, format, args...)
}

func Fprint(file *os.File, color Color, args ...interface{}) (n int, err error) {
    SetColor(file, color)
    defer UnsetColor(file)
    return fmt.Fprint(file, args...)
}

func Fprintln(file *os.File, color Color, args ...interface{}) (n int, err error) {
    SetColor(file, color)
    defer UnsetColor(file)
    return fmt.Fprintln(file, args...)
}

func Printf(color Color, format string, args ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, color, format, args...)
}

func Print(color Color, args ...interface{}) (n int, err error) {
    return Fprint(os.Stdout, color, args...)
}

func Println(color Color, args ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, color, args...)
}


func Error(args ...interface{}) (n int, err error) {
    return Fprintln(os.Stderr, errorColor, args...)
}

func Errorf(format string, args ...interface{}) (n int, err error) {
    return Fprintf(os.Stderr, errorColor, format, args...)
}

func Fatal(code int, args ...interface{}) {
    Error(args...)
    os.Exit(code)
}

func Fatalf(code int, format string, args ...interface{}) {
    Errorf(format, args...)
    os.Exit(code)
}

func Warn(args ...interface{}) (n int, err error) {
    return Fprintln(os.Stderr, warnColor, args...)
}

func Warnf(format string, args ...interface{}) (n int, err error) {
    return Fprintf(os.Stderr, warnColor, format, args...)
}

func Info(args ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, infoColor, args...)
}

func Infof(format string, args ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, infoColor, format, args...)
}

func Success(args ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, successColor, args...)
}

func Successf(format string, args ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, successColor, format, args...)
}

func Debug(args ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, debugColor, args...)
}

func Debugf(format string, args ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, debugColor, format, args...)
}

func Trace(args ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, traceColor, args...)
}

func Tracef(format string, args ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, traceColor, format, args...)
}

func Important(args ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, importantColor, args...)
}

func Importantf(format string, args ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, importantColor, format, args...)
}


func Info2(args ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, info2Color, args...)
}

func Info2f(format string, args ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, info2Color, format, args...)
}


func Colorize(color Color, format string, args ...interface{}) string {
    return fmt.Sprintf(string(color) + format + string(DefaultColor), args...)
}

func ErrorStr(format string, args ...interface{}) string {
    return Colorize(errorColor, format, args...)
}

func WarnStr(format string, args ...interface{}) string {
    return Colorize(warnColor, format, args...)
}

func InfoStr(format string, args ...interface{}) string {
    return Colorize(infoColor, format, args...)
}

func SuccessStr(format string, args ...interface{}) string {
    return Colorize(successColor, format, args...)
}

func DebugStr(format string, args ...interface{}) string {
    return Colorize(debugColor, format, args...)
}

func TraceStr(format string, args ...interface{}) string {
    return Colorize(traceColor, format, args...)
}

func ImportantStr(format string, args ...interface{}) string {
    return Colorize(importantColor, format, args...)
}

func Info2Str(format string, args ...interface{}) string {
    return Colorize(info2Color, format, args...)
}


