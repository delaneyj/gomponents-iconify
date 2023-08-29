package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSim0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M8 4h24.889L40 11.273V44H8V4Z"/><path fill="#000" stroke="#000" d="M33 26H15v10h18V26Z"/><path stroke="#000" stroke-linecap="round" d="M15 12v6m6-6v6m6-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSim0)"/>`),
		g.Group(children),
	)
}