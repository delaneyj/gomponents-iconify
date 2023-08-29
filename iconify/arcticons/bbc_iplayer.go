package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BbcIplayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.112 11.601h8.2v24.797h-8.2zm33.776 5.298L36.788 24L15.313 11.602l4.1-7.102zM19.413 43.5l-4.1-7.102L36.788 24l4.1 7.101z"/>`),
		g.Group(children),
	)
}