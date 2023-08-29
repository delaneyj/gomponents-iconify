package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Letterboxd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.55 20a7.59 7.59 0 1 0 0 8a7.59 7.59 0 0 1 0-8Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.45 20a7.59 7.59 0 0 0-12.9 0a7.64 7.64 0 0 1 0 8a7.59 7.59 0 0 0 12.9 0a7.64 7.64 0 0 1 0-8Zm6.46-3.59A7.6 7.6 0 0 0 30.45 20a7.59 7.59 0 0 1 0 8a7.59 7.59 0 1 0 6.46-11.59Z"/>`),
		g.Group(children),
	)
}