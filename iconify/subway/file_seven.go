package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M325.5 0v128h128L325.5 0zm-21.3 0H69.5v512h384V149.3H304.2V0zm64 277.3c23.5 0 42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7zm-213.4 85.4c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7zm106.7 0c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7s-19.2 42.7-42.7 42.7z"/>`),
		g.Group(children),
	)
}