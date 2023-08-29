package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetailsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19h14m1.986-2.016a2 2 0 0 0-.146-.734L13.74 4a2 2 0 0 0-3.5 0l-.821 1.417M7.95 7.951L3.14 16.25A2 2 0 0 0 4.89 19M12 3v5m0 4v7M3 3l18 18"/>`),
		g.Group(children),
	)
}