package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankNorwegian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.666 27.193c2.37-1.712 3.69-3.34 3.25-4.276c-.373-.792-1.933-.93-4.12-.507m-13.491 6.434c-1.564 1.337-2.355 2.543-2 3.298c.706 1.5 5.652.617 11.083-1.897c1.983-.917 2.356-1.13 2.356-1.13m-5.578-11.99v8.712m0-1.851l3.944-3.924m-2.688 2.675l3.1 3.087m-13.488-2.165a2.178 2.178 0 0 1-2.178 2.178h0a2.178 2.178 0 0 1-2.178-2.178v-1.416c0-1.203.975-2.178 2.178-2.178h0c1.203 0 2.178.975 2.178 2.178m0 3.594v-5.772m6.78 5.772v-3.594a2.178 2.178 0 0 0-2.178-2.178h0a2.178 2.178 0 0 0-2.179 2.178m0 3.594v-5.772M9 22.243c0-1.203.975-2.178 2.178-2.178h0c1.203 0 2.178.975 2.178 2.178v1.416a2.178 2.178 0 0 1-2.178 2.178h0A2.178 2.178 0 0 1 9 23.659m0 2.178v-8.712"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5c11.876 0 21.5 9.624 21.5 21.5S35.876 45.5 24 45.5S2.5 35.876 2.5 24S12.124 2.5 24 2.5Z"/>`),
		g.Group(children),
	)
}