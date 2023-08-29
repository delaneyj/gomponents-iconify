package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunWind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.468 10a4 4 0 1 0-5.466 5.46M2 12h1m8-9v1m0 16v1M4.6 5.6l.7.7m12.1-.7l-.7.7M5.3 17.7l-.7.7M15 13h5a2 2 0 1 0 0-4m-8 7h5.967A2 2 0 0 1 20 18a2 2 0 0 1-2 2h-.286"/>`),
		g.Group(children),
	)
}