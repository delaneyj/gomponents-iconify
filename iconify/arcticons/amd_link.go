package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmdLink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.614 33.614l9.25 9.25v-37h-37l9.25 9.25h18.5v18.5zm-18.5 0V19.55l-9.229 9.228l-.021 14.086l14.085-.022l9.228-9.228H15.114z"/>`),
		g.Group(children),
	)
}