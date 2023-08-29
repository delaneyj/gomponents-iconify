package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChinaUniversityMooc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.387 11.974v15.873c7.622 1.255 14.384 2.87 19.378 5.227V18.221a47.884 47.884 0 0 0-19.378-6.247Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.113 15.926v15.618a99.418 99.418 0 0 0-19.57 5.482V22.173c5.708-3.83 12.444-5.492 19.57-6.247Z"/>`),
		g.Group(children),
	)
}