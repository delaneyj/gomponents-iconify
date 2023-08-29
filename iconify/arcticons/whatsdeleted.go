package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whatsdeleted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.274 13.707h25.451V39.94a3.56 3.56 0 0 1-3.56 3.56h-18.33a3.56 3.56 0 0 1-3.56-3.56V13.707h0Zm19.679-6.138l-.89-3.069H17.937l-.89 3.069H9.328v6.138h29.344V7.569h-7.719z"/>`),
		g.Group(children),
	)
}