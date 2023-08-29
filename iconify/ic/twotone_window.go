package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 13h6v6h-6zm-8 0h6v6H5zm0-8h6v6H5zm8 0h6v6h-6z" opacity=".3"/><path fill="currentColor" d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm-8 16H5v-6h6v6zm0-8H5V5h6v6zm8 8h-6v-6h6v6zm0-8h-6V5h6v6z"/>`),
		g.Group(children),
	)
}