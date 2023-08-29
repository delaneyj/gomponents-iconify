package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivatTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.548 33.5l-.001-19L25.35 27.262h12.588m-27.875-6.468a6.294 6.294 0 1 1 12.587 0a5.872 5.872 0 0 1-1.843 4.45C18.26 27.479 10.063 33.5 10.063 33.5H22.65"/>`),
		g.Group(children),
	)
}