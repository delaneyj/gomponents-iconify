package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.69 16.743l.5-2a1 1 0 0 0-.97-1.243H3.78a1 1 0 0 0-.97 1.243l.5 2a1 1 0 0 0 .97.757h12.44a1 1 0 0 0 .97-.757zM17.5 11.5l.56-1.682a1 1 0 0 0-.95-1.316l-13.22.017a1 1 0 0 0-.947 1.318L3.5 11.5m14-5l.306-1.836A1 1 0 0 0 16.82 3.5H4.18a1 1 0 0 0-.986 1.164L3.5 6.5"/>`),
		g.Group(children),
	)
}