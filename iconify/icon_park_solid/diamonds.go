package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamonds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDiamonds0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M12 8h24l8 10l-20 24L4 18l8-10Z"/><path stroke="#000" d="M4 18h40M24 42l-8-24m8 24l8-24"/><path stroke="#fff" d="m8 13l-4 5l20 24l20-24l-4-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDiamonds0)"/>`),
		g.Group(children),
	)
}