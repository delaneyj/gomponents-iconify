package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RampUpArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M23 23H1v-1c14-3.228 21.5-9 21.5-9h.5v10ZM13 8.246L.5 15.5M13 8.246c-.498.29-.944.95-1.274 1.559a7.972 7.972 0 0 0-.856 2.622c-.106.712-.155 1.514.068 1.902M13 8.246c-.498.29-1.29.348-1.98.33a7.87 7.87 0 0 1-2.687-.566C7.666 7.747 6.95 7.388 6.728 7"/>`),
		g.Group(children),
	)
}