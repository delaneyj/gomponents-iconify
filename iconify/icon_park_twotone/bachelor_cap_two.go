package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BachelorCapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBachelorCapTwo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M11 21v15.039c0 .607.274 1.18.785 1.509C13.486 38.643 17.86 41 24 41s10.514-2.357 12.215-3.452c.51-.33.785-.902.785-1.51V21"/><path stroke-linecap="round" d="M43 24v8"/><path fill="#555" stroke-linecap="round" d="M5 17L24 7l19 10l-19 10L5 17Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBachelorCapTwo0)"/>`),
		g.Group(children),
	)
}