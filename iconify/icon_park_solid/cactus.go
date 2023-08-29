package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cactus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCactus0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 43h32M6 16c0 5.5.5 14 10 14m26-16c0 5 0 20-10 20"/><path fill="#fff" d="M24 5a8 8 0 0 0-8 8v30h16V13a8 8 0 0 0-8-8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCactus0)"/>`),
		g.Group(children),
	)
}