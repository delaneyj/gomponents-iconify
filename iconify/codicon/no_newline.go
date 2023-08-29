package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoNewline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.333 5.506a3 3 0 1 1 3.334 4.989a3 3 0 0 1-3.334-4.99zm2.677.777A1.986 1.986 0 0 0 2 8.009c.004.353.102.698.283 1.001L5.01 6.283zM2.99 9.717A1.986 1.986 0 0 0 6 7.991a1.988 1.988 0 0 0-.283-1.001L2.99 9.717zM14 5v1.984a.5.5 0 0 1-.5.5H9.367L11 5.851l-.707-.707l-2.121 2.121l-.423.423v.568l2.544 2.544l.707-.707l-1.61-1.609h4.11a1.5 1.5 0 0 0 1.5-1.5V5h-1z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}