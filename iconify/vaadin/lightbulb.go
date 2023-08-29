package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a5 5 0 0 0-5 5a4.8 4.8 0 0 0 2.182 3.989L6 11a.5.5 0 0 0 0 1a.5.5 0 0 0 0 1a.5.5 0 0 0 0 1a.5.5 0 0 0 0 1h.41c.342.55.915.929 1.581.999a2.122 2.122 0 0 0 1.594-.99L10 15a.5.5 0 0 0 0-1a.5.5 0 0 0 0-1a.5.5 0 0 0 0-1a.5.5 0 0 0 0-1l.8-2A4.805 4.805 0 0 0 13 5.002A5.001 5.001 0 0 0 8 0zm2.25 8.21l-.25.17l-.11.29L9 10.81a.292.292 0 0 1-.27.19H7.22a.29.29 0 0 1-.219-.188L6.13 8.67L6 8.38l-.25-.18A3.914 3.914 0 0 1 4 5.003A4 4 0 1 1 12 5a3.905 3.905 0 0 1-1.736 3.201z"/><path fill="currentColor" d="M10.29 3a3.153 3.153 0 0 0-2.289-1L8 3a2.133 2.133 0 0 1 1.5.62c.278.386.459.858.499 1.37L11 4.999a3.785 3.785 0 0 0-.718-2.011z"/>`),
		g.Group(children),
	)
}