package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.981 24.433A3.955 3.955 0 0 0 26 25a4.005 4.005 0 0 0 4-4a3.943 3.943 0 0 0-.15-1.023l-2.436 2.437a2 2 0 1 1-2.828-2.828l2.437-2.437A3.95 3.95 0 0 0 26 17a4.005 4.005 0 0 0-4 4a3.954 3.954 0 0 0 .567 2.02L17 28.585L18.414 30Z"/><path fill="currentColor" fill-rule="evenodd" d="M26 4h-4V2h-2v2h-8V2h-2v2H6a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h8v-2H6V6h4v2h2V6h8v2h2V6h4v8h2V6a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}