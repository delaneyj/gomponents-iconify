package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M160 0c17.7 0 32 14.3 32 32v32h128V32c0-17.7 14.3-32 32-32s32 14.3 32 32v32h48c26.5 0 48 21.5 48 48v48H32v-48c0-26.5 21.5-48 48-48h48V32c0-17.7 14.3-32 32-32zM32 192h448v272c0 26.5-21.5 48-48 48H80c-26.5 0-48-21.5-48-48V192zm312 184c13.3 0 24-10.7 24-24s-10.7-24-24-24H168c-13.3 0-24 10.7-24 24s10.7 24 24 24h176z"/>`),
		g.Group(children),
	)
}