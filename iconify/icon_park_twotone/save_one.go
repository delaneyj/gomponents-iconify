package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSaveOne0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M39.3 6H8.7A2.7 2.7 0 0 0 6 8.7v30.6A2.7 2.7 0 0 0 8.7 42h30.6a2.7 2.7 0 0 0 2.7-2.7V8.7A2.7 2.7 0 0 0 39.3 6Z"/><path fill="#555" stroke-linejoin="round" d="M32 6v18H15V6h17Z"/><path stroke-linecap="round" d="M26 13v4M10.997 6H36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSaveOne0)"/>`),
		g.Group(children),
	)
}