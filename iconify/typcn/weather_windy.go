package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherWindy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 5a1 1 0 1 0 0 2a1 1 0 0 1 0 2H8a1 1 0 1 0 0 2h6a1 1 0 0 1 0 2H7.6c-1.654 0-3 1.346-3 3s1.346 3 3 3a1 1 0 1 0 0-2a1 1 0 0 1 0-2H14c1.654 0 3-1.346 3-3c0-.353-.072-.686-.185-1H19c1.654 0 3-1.346 3-3s-1.346-3-3-3z"/>`),
		g.Group(children),
	)
}