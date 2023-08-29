package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonumentTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMonumentTwo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4 44h40"/><path fill="#fff" d="m18 44l2-36.842L28 4l2 40H18Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMonumentTwo0)"/>`),
		g.Group(children),
	)
}