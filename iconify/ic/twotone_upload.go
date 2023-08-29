package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.83 8H11v6h2V8h1.17L12 5.83z" opacity=".3"/><path fill="currentColor" d="m12 3l-7 7h4v6h6v-6h4l-7-7zm1 5v6h-2V8H9.83L12 5.83L14.17 8H13zM5 18h14v2H5z"/>`),
		g.Group(children),
	)
}