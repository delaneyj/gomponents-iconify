package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 24c0-8.284 6.716-15 15-15c8.284 0 15 6.716 15 15c0 8.284-6.716 15-15 15c-8.284 0-15-6.716-15-15Zm15-13c-7.18 0-13 5.82-13 13s5.82 13 13 13s13-5.82 13-13s-5.82-13-13-13Zm0 4a9 9 0 1 0 0 18a9 9 0 0 0 0-18Zm-11 9c0-6.075 4.925-11 11-11s11 4.925 11 11s-4.925 11-11 11s-11-4.925-11-11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}