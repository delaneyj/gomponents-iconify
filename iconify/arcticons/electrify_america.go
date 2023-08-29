package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectrifyAmerica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.223 7.109l11.982-2.733l-4.875 15.217l7.29-1.731l-19.956 25.514l4.01-14.58l-5.468 1.185l7.017-22.872z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.602 29.056c.81.43 1.866.668 3.187.669c3.639.002 5.73-2.784 5.73-2.784m.375-9.107c-2.443-2.958-9.216-.826-10.872 4.288c-.973 3.003-.665 5.51 1.275 6.756"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.902 22.805c7.18.978 12.51-.302 11.72-3.519"/>`),
		g.Group(children),
	)
}