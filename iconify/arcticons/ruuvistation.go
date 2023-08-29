package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruuvistation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.287 10.205a13.795 13.795 0 0 0-5.037 1.002a6.506 6.506 0 0 1-4.554 11.156h0a6.509 6.509 0 0 1-3.786-1.215a13.794 13.794 0 1 0 13.496-10.943h-.119Z"/>`),
		g.Group(children),
	)
}