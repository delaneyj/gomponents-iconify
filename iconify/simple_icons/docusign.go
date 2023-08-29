package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Docusign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.517 3.31h4.966v6.621h3.31L12 16.552L6.207 9.931h3.31V3.31zM0 19.034h24v1.655H0v-1.655z"/>`),
		g.Group(children),
	)
}