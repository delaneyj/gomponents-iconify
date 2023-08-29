package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortNumeric(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 2h-2v2H9v2h2V4h2v2h2V4h-2V2zM2 8h4v8H4v-6H2V8zm6 0h6v5h-4v1h4v2H8v-5h4v-1H8V8zm12 0h-4v2h4v1h-4v2h4v1h-4v2h6V8h-2zm-9 10v2H9v-2h2zm2 2h-2v2h2v-2zm0 0v-2h2v2h-2z"/>`),
		g.Group(children),
	)
}