package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactoryBuilding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFactoryBuilding0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 44V4h8v16l16-8v8l16-8v32H4Z"/><path fill="#000" stroke="#000" d="M12 28h8v8h-8zm16 0h8v8h-8z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFactoryBuilding0)"/>`),
		g.Group(children),
	)
}