package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 10a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v10H12V10Zm32 25c-1.108 1.333-2.375 5-7.6 5c-2.737 0-6.456-1.684-11.4-3"/><path d="M4 35s6-9 8-15h24c2 6 8 15 8 15c-6-4-25 5-32 5c-5.5 0-6.833-3.667-8-5Z"/></g>`),
		g.Group(children),
	)
}