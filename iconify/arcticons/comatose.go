package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comatose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22 13.3c0-3.9 1.1-7.6 2.9-10.8H24C12.1 2.5 2.5 12.1 2.5 24S12.1 45.5 24 45.5c8 0 14.9-4.3 18.6-10.7C31.2 34.3 22 24.9 22 13.3Z"/>`),
		g.Group(children),
	)
}