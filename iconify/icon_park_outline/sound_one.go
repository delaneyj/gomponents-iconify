package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="4" d="M9 7a3 3 0 0 1 3-3h24a3 3 0 0 1 3 3v34a3 3 0 0 1-3 3H12a3 3 0 0 1-3-3V7Z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M24 29a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><rect width="4" height="4" x="30" y="8" fill="currentColor" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M27 36h2m-10 0h2"/></g>`),
		g.Group(children),
	)
}