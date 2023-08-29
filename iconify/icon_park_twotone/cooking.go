package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cooking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCooking0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M6 42h36M6 36h36"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M9 25c0-8.284 6.716-15 15-15c8.284 0 15 6.716 15 15v11H9V25Z"/><path stroke-linecap="round" d="M17 25v4"/><path d="M28 10V8a4 4 0 0 0-8 0v2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCooking0)"/>`),
		g.Group(children),
	)
}