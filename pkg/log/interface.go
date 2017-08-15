package log

type Lvl int

const (
    LvlCrit Lvl = iota
    LvlError
    LvlWarn
    LvlInfo
    LvlDebug
)
