package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsKitchen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 3h8l-1 9H5zm3 15h2v3H7zM20 3v12h-5c-.023-3.681.184-7.406 5-12zm0 12v6h-1v-3M8 12v6"/>`),
		g.Group(children),
	)
}