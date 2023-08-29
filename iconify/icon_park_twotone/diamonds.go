package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamonds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDiamonds0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M12 8h24l8 10l-20 24L4 18l8-10Z"/><path d="M4 18h40M24 42l-8-24m8 24l8-24M8 13l-4 5l20 24l20-24l-4-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDiamonds0)"/>`),
		g.Group(children),
	)
}