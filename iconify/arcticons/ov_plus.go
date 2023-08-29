package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.218 8.018a.972.972 0 0 0-.866-.532H6.473a.972.972 0 0 0-.867 1.412L13.27 24L5.606 39.102a.972.972 0 0 0 .867 1.412h9.878a.972.972 0 0 0 .867-.532L25.328 24l-8.11-15.981Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.39 8.018a.972.972 0 0 0-.866-.532h-9.879a.972.972 0 0 0-.866 1.412L30.442 24l-7.663 15.102a.972.972 0 0 0 .866 1.412h9.879a.972.972 0 0 0 .867-.532L42.5 24L34.39 8.019Z"/>`),
		g.Group(children),
	)
}