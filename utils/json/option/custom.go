package option

import (
	"strconv"

	"github.com/diamondburned/arikawa/discord"
)

// ================================ Seconds ================================

// Seconds is the option type for discord.Seconds.
type Seconds = *discord.Seconds

// ZeroSeconds are 0 Seconds.
var ZeroSeconds = NewSeconds(0)

// NewString creates a new Seconds with the value of the passed discord.Seconds.
func NewSeconds(s discord.Seconds) Seconds { return &s }

// ================================ Color ================================

// Color is the option type for discord.Color.
type Color = *discord.Color

// NewString creates a new Color with the value of the passed discord.Color.
func NewColor(s discord.Color) Color { return &s }

// ================================ NullableColor ================================

// Nullable is a nullable version of discord.Color.
type NullableColor = *nullableColor

type nullableColor struct {
	Val  discord.Color
	Init bool
}

// NullColor serializes to JSON null.
var NullColor = &nullableColor{}

// NewNullableColor creates a new non-null NullableColor using the value of the
// passed discord.Color.
func NewNullableColor(v discord.Color) NullableColor {
	return &nullableColor{
		Val:  v,
		Init: true,
	}
}

func (i nullableColor) MarshalJSON() ([]byte, error) {
	if !i.Init {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatUint(uint64(i.Val), 10)), nil
}

func (i *nullableColor) UnmarshalJSON(json []byte) error {
	s := string(json)

	if s == "null" {
		i.Init = false
		return nil
	}

	v, err := strconv.ParseUint(s, 10, 32)

	i.Val = discord.Color(v)

	return err
}
