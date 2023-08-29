package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTExpandLeft0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Z"/><path stroke-linecap="round" d="M32 6v36M16 20l4 4l-4 4M26 6h12M26 42h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTExpandLeft0)"/>`),
		g.Group(children),
	)
}