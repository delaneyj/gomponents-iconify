package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Won(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m4 6l1.813 6H4v2h2.406L7 16H4v2h3.594L10 26h2l2-8h4l2 8h2l2.406-8H28v-2h-3l.594-2H28v-2h-1.813L28 6h-2l-1.813 6H18.5L17 6h-2l-1.5 6H7.812L6 6H4zm12 4l.5 2h-1l.5-2zm-7.594 4H13l-.5 2H9l-.594-2zM15 14h2l.5 2h-3l.5-2zm4 0h4.594L23 16h-3.5l-.5-2zm-9.406 4H12l-1.094 4.375L9.594 18zM20 18h2.406l-1.312 4.375L20 18z"/>`),
		g.Group(children),
	)
}