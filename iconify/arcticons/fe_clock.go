package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 1 2.5 24A21.5 21.5 0 0 1 24 2.5Z"/><circle cx="24" cy="24" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.232 22.232l-7.113-7.113m10.649 7.118l10.227-10.251M25.44 41.064h5.92m-5.92-11.84h5.92m-5.92 5.92h3.86m-3.86-5.92v11.84m-8.8-11.84h5.92m-5.92 5.92h3.848m-3.848-5.92v11.84"/>`),
		g.Group(children),
	)
}