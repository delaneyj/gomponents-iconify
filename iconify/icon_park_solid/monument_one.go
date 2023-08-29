package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonumentOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMonumentOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M14 38h20v6H14z"/><path fill="#fff" d="m18 38l2-29l4-5l4 5l2 29H18Z"/><path stroke-linecap="round" d="M4 44h40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMonumentOne0)"/>`),
		g.Group(children),
	)
}