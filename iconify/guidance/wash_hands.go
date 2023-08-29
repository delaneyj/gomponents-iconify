package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashHands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 23.5h21.5V23a3 3 0 0 0-3-3H12h11.5v-.5a3 3 0 0 0-3-3H12M0 13a3.5 3.5 0 0 0 3.5-3.5h11v.5a3 3 0 0 1-3 3H8h10.5a3 3 0 0 1 3 3v.67m2-10.17c0 1.638-1.343 3.004-3 3.004s-3-1.366-3-3.004c0-.158.027-.324.072-.49c.284-1.07 1.045-1.94 1.749-2.793L19.5 3c.99-1.363 1-2.5 1-3c0 .5.01 1.637 1 3l.18.217c.703.853 1.464 1.723 1.748 2.792c.045.167.072.333.072.491Z"/>`),
		g.Group(children),
	)
}