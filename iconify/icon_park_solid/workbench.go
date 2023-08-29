package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Workbench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWorkbench0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M12 33H4V7h40v26H12Z"/><path stroke="#000" stroke-linecap="round" d="M16 22v4"/><path stroke="#fff" stroke-linecap="round" d="M24 33v6"/><path stroke="#000" stroke-linecap="round" d="M24 18v8m8-12v12"/><path stroke="#fff" stroke-linecap="round" d="M12 41h24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWorkbench0)"/>`),
		g.Group(children),
	)
}