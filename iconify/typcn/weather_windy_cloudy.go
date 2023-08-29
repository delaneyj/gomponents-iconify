package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherWindyCloudy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.798 15.75a.978.978 0 0 1-.4-.084A4.004 4.004 0 0 1 2 12a4.007 4.007 0 0 1 3.001-3.874L5 8c0-3.309 2.691-6 6-6a5.974 5.974 0 0 1 5.902 5.001a1 1 0 0 1-.82 1.152a1.007 1.007 0 0 1-1.152-.82A3.98 3.98 0 0 0 11 4a4.004 4.004 0 0 0-3.919 4.812l.259 1.27l-1.431-.088C4.897 10 4 10.897 4 12c0 .795.471 1.515 1.2 1.834a1 1 0 0 1-.402 1.916zM19 7a1 1 0 1 0 0 2a1 1 0 0 1 0 2H9.4a1 1 0 1 0 0 2H14a1 1 0 0 1 0 2H9c-1.654 0-3 1.346-3 3s1.346 3 3 3a1 1 0 1 0 0-2a1 1 0 0 1 0-2h5c1.654 0 3-1.346 3-3c0-.353-.072-.686-.185-1H19c1.654 0 3-1.346 3-3s-1.346-3-3-3z"/>`),
		g.Group(children),
	)
}