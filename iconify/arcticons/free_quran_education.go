package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreeQuranEducation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 12.263h9.915m18.563 0H43.5m-6.074 10.301h5.989M4.577 36.037V22.564h13.302m11.076 9.566h14.287"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.736 19.028a10.736 10.736 0 1 1 0 7.344"/>`),
		g.Group(children),
	)
}