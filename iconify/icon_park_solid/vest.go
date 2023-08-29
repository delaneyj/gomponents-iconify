package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSVest0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 4h-5l-7 20v20h16V19l-4-6V4ZM12 4h5l7 20v20H8V19l4-6V4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSVest0)"/>`),
		g.Group(children),
	)
}