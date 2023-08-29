package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backpack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBackpack0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M19 9.556V4h-6v10m16-4.444V4h6v10"/><path fill="#fff" stroke="#fff" d="M11 20c0-5.523 4.477-10 10-10h6c5.523 0 10 4.477 10 10v20a4 4 0 0 1-4 4H15a4 4 0 0 1-4-4V20Z"/><path stroke="#fff" d="M11 29H5v10h6m26-10h6v10h-6"/><path stroke="#000" d="M28 23v4m-11-4h14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBackpack0)"/>`),
		g.Group(children),
	)
}