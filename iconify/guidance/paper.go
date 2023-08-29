package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M13.5 1.5v7h7m-16-7v21h16V8l-.282-.126a12 12 0 0 1-6.092-6.092L14 1.5H4.5Z"/>`),
		g.Group(children),
	)
}