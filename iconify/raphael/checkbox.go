package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 27.5H6A1.5 1.5 0 0 1 4.5 26V6c0-.83.67-1.5 1.5-1.5h20c.828 0 1.5.67 1.5 1.5v20a1.5 1.5 0 0 1-1.5 1.5zm-18.5-3h17v-17h-17v17z"/>`),
		g.Group(children),
	)
}