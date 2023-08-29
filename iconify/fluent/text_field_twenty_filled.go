package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFieldTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6Zm4-.5v1a.5.5 0 0 0 1 0V6h2.5v8H9a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-.5V6H13v.5a.5.5 0 0 0 1 0v-1a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.5.5Z"/>`),
		g.Group(children),
	)
}