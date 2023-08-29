package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xingfuli(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSXingfuli0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path fill="#000" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m32 17l-12-4v19l6 3V19l6-2Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSXingfuli0)"/>`),
		g.Group(children),
	)
}