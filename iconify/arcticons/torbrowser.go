package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torbrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 19a5 5 0 0 1 0 10"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 13.72a10.28 10.28 0 0 1 0 20.56"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.88a16.12 16.12 0 0 1 0 32.24"/>`),
		g.Group(children),
	)
}