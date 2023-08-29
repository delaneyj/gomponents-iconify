package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDisk0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M41 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Z"/><path fill="#000" stroke="#000" stroke-linejoin="round" d="M34 4v18H15V4h19Z"/><path stroke="#000" stroke-linecap="round" d="M29 11v4"/><path stroke="#fff" stroke-linecap="round" d="M11.997 4h25.001"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDisk0)"/>`),
		g.Group(children),
	)
}