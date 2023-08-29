package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airpods(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAirpods0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M36 4a9 9 0 0 0-9 9v31h6V21.488A8.987 8.987 0 0 0 36 22a9 9 0 0 0 8-4.873V8.873A9 9 0 0 0 36 4ZM12 4a9 9 0 0 1 9 9v31h-6V21.488A8.987 8.987 0 0 1 12 22a9 9 0 0 1-8-4.873V8.873A9 9 0 0 1 12 4Z"/><path stroke="#000" d="M15 13h-1m19 0h1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAirpods0)"/>`),
		g.Group(children),
	)
}