package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluentThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.553 2.106a1 1 0 0 1 .894 0l8 4a1 1 0 0 1 0 1.788L19.237 11l6.21 3.106a1 1 0 0 1 0 1.788L18 19.618V29a1 1 0 0 1-1.49.872l-8-4.5A1 1 0 0 1 8 24.5V7a1 1 0 0 1 .553-.894l8-4ZM10 7.618v16.297l6 3.375V19a1 1 0 0 1 .553-.894L22.763 15l-6.21-3.106a1 1 0 0 1 0-1.788L22.763 7L17 4.118l-7 3.5Z"/>`),
		g.Group(children),
	)
}