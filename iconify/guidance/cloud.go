package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10 9.362C11 6.5 13.5 4.5 16.536 4.5c3.964 0 6.964 3.154 6.964 7a6 6 0 0 1-6 6H.5v-.536l.367-.446c1.471-1.79 1.984-4.325 3.558-6.025A4.867 4.867 0 0 1 8 8.93a4.821 4.821 0 0 1 4.821 4.821"/>`),
		g.Group(children),
	)
}