package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 32v32H80c-26.5 0-48 21.5-48 48v48h448v-48c0-26.5-21.5-48-48-48h-48V32c0-17.7-14.3-32-32-32s-32 14.3-32 32v32H192V32c0-17.7-14.3-32-32-32s-32 14.3-32 32zm352 160H32v272c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V192zm-224 56c13.3 0 24 10.7 24 24v56h56c13.3 0 24 10.7 24 24s-10.7 24-24 24h-56v56c0 13.3-10.7 24-24 24s-24-10.7-24-24v-56h-56c-13.3 0-24-10.7-24-24s10.7-24 24-24h56v-56c0-13.3 10.7-24 24-24z"/>`),
		g.Group(children),
	)
}