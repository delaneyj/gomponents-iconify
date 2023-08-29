package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IUtensils(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M28 4h-9v40h9v-4h-4v-2h4v-3h-2v-2h2v-3h-2v-2h2v-3h-4v-2h4v-3h-2v-2h2v-3h-2v-2h2v-3h-4V8h4V4Zm5 24a2 2 0 0 0-2 2v5h7v-5a2 2 0 0 0-2-2h-3Zm5 9h-7v5a2 2 0 0 0 2 2h3a2 2 0 0 0 2-2v-5Z"/><path d="M33 30v3h3v-3h-3Zm5 0a2 2 0 0 0-2-2h-3a2 2 0 0 0-2 2v5h7v-5Zm-7 7h7v5a2 2 0 0 1-2 2h-3a2 2 0 0 1-2-2v-5Zm5 2h-3v3h3v-3ZM13 11a3 3 0 0 0-3 3v3h6v-3a3 3 0 0 0-3-3Zm3 8h-6v21l3 4l3-4V19Z"/></g>`),
		g.Group(children),
	)
}