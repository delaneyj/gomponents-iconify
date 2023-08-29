package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BringToFrontOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBringToFrontOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M14 21v13h13m-6-20h13v13"/><path fill="#fff" d="M5 21V5h16v16H5Zm22 22V27h16v16H27Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBringToFrontOne0)"/>`),
		g.Group(children),
	)
}