package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Workbench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWorkbench0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M12 33H4V7h40v26H12Z"/><path stroke-linecap="round" d="M16 22v4m8 7v6m0-21v8m8-12v12M12 41h24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWorkbench0)"/>`),
		g.Group(children),
	)
}