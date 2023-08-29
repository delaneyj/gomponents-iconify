package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xingfuli(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTXingfuli0"><g fill="#555" stroke="#fff" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m32 17l-12-4v19l6 3V19l6-2Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTXingfuli0)"/>`),
		g.Group(children),
	)
}