package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamWiredStandOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q-1.25 0-2.125.863T9 20v1h6v-1q0-1.275-.863-2.138T12 17ZM8 7.75q0 1.675 1.125 2.9t2.8 1.425l4.1.425V3l-4.1.4q-1.675.175-2.8 1.413T8 7.75ZM17 23H7v-3q0-2.1 1.463-3.55T12 15q.275 0 .525.025t.5.075l.55-.85l-1.85-.175q-2.425-.25-4.075-2.063T6 7.75q0-2.475 1.638-4.288T11.724 1.4l6.3-.6v13.875l-2.2-.2l-.95 1.425q.975.675 1.55 1.75T17 20v3Zm-5-2Zm4.025-13.25Z"/>`),
		g.Group(children),
	)
}