package cli

import (
    "strconv"
    "fmt"
    "time"
)

type optionType string

const (
    flag     optionType = "flag"
    value    optionType = "value"
    integer  optionType = "number"
    float    optionType = "floating-point number"
    duration optionType = "duration"
    path     optionType = "path"

    indentSize = 4

    longPrefix  = "--"
    shortPrefix = "-"
)

type Option struct {
    long     string
    desc     string
    short    string
    used     bool
    required bool
    argType  optionType
    defVal   *string
    setter   func(string) error
}

func (o *Option) expects() string {
    switch o.argType {
    case flag:
        return ""
    default:
        return "<" + string(o.argType) + ">"
    }
}

func (o *Option) defaultValue() string {
    if o.defVal != nil {
        return fmt.Sprintf("(default %s)", *o.defVal)
    }
    return ""
}

func (o *Option) trigger() string {
    return fmt.Sprintf("%s, %s %s", o.long, o.short, o.expects())
}

func (o *Option) description() string {
    return fmt.Sprintf("%s %s", o.desc, o.defaultValue())
}

func (o *Option) set(value string) error {
    if err := o.setter(value); err != nil {
        return err
    }
    o.used = true
    return nil
}

func notNil(value interface{}, long string, short byte) interface{} {
    if value == nil {
        panic("Value for option " + longPrefix + long + "|" + shortPrefix + string(short) + " can't be nil")
    }
    return value
}

func Required(option *Option) *Option {
    option.required = true
    return option
}

func newOption(optType optionType, long string, short byte, description string, setter func(string) error, defaultValue *string) *Option {
    long = Escape(long)
    option := &Option{
        long:    longPrefix + long,
        short:   shortPrefix + string(short),
        desc:    Sentence(description),
        argType: optType,
        defVal:  defaultValue,
        setter:  setter}
    if defaultValue != nil {
        option.set(*defaultValue)
    }
    return option
}

func FlagOptFunc(handler func() error, long string, short byte, description string) *Option {
    return newOption(flag, long, short, description, func(string) error {
        return handler()
    }, nil)
}

func FlagOpt(value *bool, long string, short byte, description string) *Option {
    return FlagOptFunc(func() error {
        val := true
        *value = val
        return nil
    }, long, short, description)
}

func IntOptFunc(handler func(int64) error, long string, short byte, description string, defaults ...int64) *Option {
    var defVal *string
    if len(defaults) > 0 {
        converted := strconv.FormatInt(defaults[0], 10)
        defVal = &converted
    }
    return newOption(integer, long, short, description, func(val string) error {
        number, err := strconv.ParseInt(val, 10, 64)
        if err != nil {
            return err
        }
        return handler(number)
    }, defVal)
}

func IntOpt(value **int64, long string, short byte, description string, defaults ...int64) *Option {
    return IntOptFunc(func(number int64) error {
        *value = &number
        return nil
    }, long, short, description, defaults...)
}

func RequiredIntOptFunc(handler func(int64) error, long string, short byte, description string, defaults ...int64) *Option {
    return Required(IntOptFunc(handler, long, short, description, defaults...))
}

func RequiredIntOpt(value *int64, long string, short byte, description string, defaults ...int64) *Option {
    notNil(value, long, short)
    return Required(IntOptFunc(func(number int64) error {
        *value = number
        return nil
    }, long, short, description, defaults...))
}

func FloatOptFunc(handler func(float64) error, long string, short byte, description string, defaults ...float64) *Option {
    var defVal *string
    if len(defaults) > 0 {
        converted := strconv.FormatFloat(defaults[0], 'f', 6, 64)
        defVal = &converted
    }
    return newOption(float, long, short, description, func(val string) error {
        number, err := strconv.ParseFloat(val, 64)
        if err != nil {
            return err
        }
        return handler(number)
    }, defVal)
}

func FloatOpt(value **float64, long string, short byte, description string, defaults ...float64) *Option {
    return FloatOptFunc(func(number float64) error {
        *value = &number
        return nil
    }, long, short, description, defaults...)
}

func RequiredFloatOptFunc(handler func(float64) error, long string, short byte, description string, defaults ...float64) *Option {
    return Required(FloatOptFunc(handler, long, short, description, defaults...))
}

func RequiredFloatOpt(value *float64, long string, short byte, description string, defaults ...float64) *Option {
    notNil(value, long, short)
    return Required(FloatOptFunc(func(number float64) error {
        *value = number
        return nil
    }, long, short, description, defaults...))
}

func StringOptFunc(handler func(string) error, long string, short byte, description string, defaults ...string) *Option {
    var defVal *string
    if len(defaults) > 0 {
        defVal = &defaults[0]
    }
    return newOption(value, long, short, description, func(val string) error {
        return handler(val)
    }, defVal)
}

func StringOpt(value **string, long string, short byte, description string, defaults ...string) *Option {
    return StringOptFunc(func(val string) error {
        *value = &val
        return nil
    }, long, short, description, defaults...)
}

func RequiredStringOptFunc(handler func(string) error, long string, short byte, description string, defaults ...string) *Option {
    return Required(StringOptFunc(handler, long, short, description, defaults...))
}

func RequiredStringOpt(value *string, long string, short byte, description string, defaults ...string) *Option {
    notNil(value, long, short)
    return Required(StringOptFunc(func(val string) error {
        *value = val
        return nil
    }, long, short, description, defaults...))
}

func DurationOptFunc(handler func(time.Duration) error, long string, short byte, description string, defaults ...time.Duration) *Option {
    var defVal *string
    if len(defaults) > 0 {
        converted := fmt.Sprintf("%v", defaults[0])
        defVal = &converted
    }
    return newOption(duration, long, short, description, func(val string) error {
        duration, err := time.ParseDuration(val)
        if err != nil {
            return err
        }
        return handler(duration)
    }, defVal)
}

func DurationOpt(value **time.Duration, long string, short byte, description string, defaults ...time.Duration) *Option {
    return DurationOptFunc(func(val time.Duration) error {
        *value = &val
        return nil
    }, long, short, description, defaults...)
}

func RequiredDurationOptFunc(handler func(time.Duration) error, long string, short byte, description string, defaults ...time.Duration) *Option {
    return Required(DurationOptFunc(handler, long, short, description, defaults...))
}

func RequiredDurationOpt(value *time.Duration, long string, short byte, description string, defaults ...time.Duration) *Option {
    notNil(value, long, short)
    return Required(DurationOptFunc(func(val time.Duration) error {
        *value = val
        return nil
    }, long, short, description, defaults...))
}

func PathOptFunc(handler func(string) error, long string, short byte, description string, defaults ...string) *Option {
    var defVal *string
    if len(defaults) > 0 {
        defVal = &defaults[0]
    }
    return newOption(path, long, short, description, func(val string) error {
        return handler(val)
    }, defVal)
}

func PathOpt(value **string, long string, short byte, description string, defaults ...string) *Option {
    return PathOptFunc(func(val string) error {
        *value = &val
        return nil
    }, long, short, description, defaults...)
}

func RequiredPathOptFunc(handler func(string) error, long string, short byte, description string, defaults ...string) *Option {
    return Required(PathOptFunc(handler, long, short, description, defaults...))
}

func RequiredPathOpt(value *string, long string, short byte, description string, defaults ...string) *Option {
    notNil(value, long, short)
    return Required(PathOptFunc(func(val string) error {
        *value = val
        return nil
    }, long, short, description, defaults...))
}
