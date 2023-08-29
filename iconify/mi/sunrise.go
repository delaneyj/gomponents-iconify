package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.707 2.293a1 1 0 0 0-1.414 0l-2 2a1 1 0 0 0 1.414 1.414L11 5.414v2.729c0 .473.448.857 1 .857s1-.384 1-.857V5.414l.293.293a1 1 0 1 0 1.414-1.414l-2-2zM5.636 11.05A1 1 0 0 0 7.05 9.636l-.707-.707a1 1 0 0 0-1.414 1.414l.707.707zm13.435-.707l-.707.707a1 1 0 0 1-1.414-1.414l.707-.707a1 1 0 0 1 1.414 1.414zM8 16a4 4 0 1 1 6.646 3H9.354A3.988 3.988 0 0 1 8 16zm.97 5H21a1 1 0 1 0 0-2h-3.803a6 6 0 1 0-10.394 0H3a1 1 0 1 0 0 2h5.97zM4 17a1 1 0 1 0 0-2H3a1 1 0 1 0 0 2h1zm18-1a1 1 0 0 1-1 1h-1a1 1 0 1 1 0-2h1a1 1 0 0 1 1 1z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}