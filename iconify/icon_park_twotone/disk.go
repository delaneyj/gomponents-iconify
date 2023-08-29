package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDisk0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M41 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Z"/><path fill="#555" stroke-linejoin="round" d="M34 4v18H15V4h19Z"/><path stroke-linecap="round" d="M29 11v4M11.997 4h25.001"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDisk0)"/>`),
		g.Group(children),
	)
}