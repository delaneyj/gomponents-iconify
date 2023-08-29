package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm24-6V6h-2v10h6v-2h-4zm-4-8h-2l-1.5 4L19 6h-2v10h2V9l1.5 4L22 9v7h2V6zM9 8h2v8h2V8h2V6H9v2zM5 6v4H2V6H0v10h2v-4h3v4h2V6H5z"/>`),
		g.Group(children),
	)
}