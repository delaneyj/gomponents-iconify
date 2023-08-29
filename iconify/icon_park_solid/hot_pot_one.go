package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotPotOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHotPotOne0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M24 36c12 0 17-8.059 17-18H7c0 9.941 5 18 17 18Z"/><path stroke="#fff" stroke-linejoin="round" d="m17 35l-3 9h20l-3-9m-2-17L27.889 4H20.11L19 18"/><path stroke="#000" d="M15 25s.07 1.07 1 2c.93.93 2 1 2 1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHotPotOne0)"/>`),
		g.Group(children),
	)
}