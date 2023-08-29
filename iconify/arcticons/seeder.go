package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seeder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="14" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.034 30.65C20.137 31.576 21.327 32 24 32h1.355a3.996 3.996 0 0 0 3.99-4h0a3.996 3.996 0 0 0-3.99-4h-2.71a3.996 3.996 0 0 1-3.99-4h0a3.996 3.996 0 0 1 3.99-4H24c2.673 0 3.863.424 4.966 1.35M24 2.5l3.963 3.75L24 10m0 35.5l-3.963-3.75L24 38"/>`),
		g.Group(children),
	)
}