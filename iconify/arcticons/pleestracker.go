package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pleestracker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.69 30.66a1 1 0 0 0-1.19-.08A15.65 15.65 0 0 1 25.33 4.41a1 1 0 0 0-.54-1.87h-.27L24 2.5a21.5 21.5 0 1 0 20 29.32a1 1 0 0 0-.31-1.16Zm-30.17 6.5a6.79 6.79 0 1 1 6.79-6.79a6.8 6.8 0 0 1-6.79 6.79Z"/><circle cx="13.52" cy="30.37" r="1.18" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.52 25.53v3.66m1.19 1.18h3.65"/>`),
		g.Group(children),
	)
}