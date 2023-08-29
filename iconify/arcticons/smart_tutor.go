package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartTutor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.578 34.811c-.45 4.136-2.05 6.498-3.572 5.277s-2.393-5.564-1.944-9.702s2.046-6.505 3.57-5.289s2.395 5.555 1.948 9.694m29.631.467c-.37 3.4-1.516 5.6-2.757 5.599a1.302 1.302 0 0 1-.815-.322c-1.522-1.221-2.393-5.564-1.944-9.702s2.046-6.505 3.57-5.289s2.395 5.555 1.948 9.694"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.945 32.594c.183-6.495.113-16.125 2.127-19.792c5.845-10.649 25.404-10.123 30.525-.473c2.196 4.138 2.833 13.937 2.732 20.712M25.08 38.432l1.91 2.571l11.463-.146"/>`),
		g.Group(children),
	)
}