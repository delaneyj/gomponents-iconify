package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Html(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 19v-8h-2v10h6v-2h-4zm-4-8h-2l-1.5 4l-1.5-4h-2v10h2v-7l1.5 4l1.5-4v7h2V11zM9 13h2v8h2v-8h2v-2H9v2zm-4-2v4H2v-4H0v10h2v-4h3v4h2V11H5z"/>`),
		g.Group(children),
	)
}