package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSkates0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M5 5h26v8H5V5Z"/><path fill="#555" d="M9 36V13h18v10c14 0 14 9 14 13H9Z"/><path d="M7 36h36v6H7z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSkates0)"/>`),
		g.Group(children),
	)
}