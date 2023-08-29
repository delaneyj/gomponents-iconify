package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upstox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.23 22.579v2.765a4.73 4.73 0 0 0 4.732 4.731h0a4.73 4.73 0 0 0 4.73-4.73v-7.807m0 7.806v4.731m5.346-4.731a4.73 4.73 0 0 0 4.731 4.731h0a4.73 4.73 0 0 0 4.731-4.73v-3.076a4.73 4.73 0 0 0-4.73-4.73h0a4.73 4.73 0 0 0-4.732 4.73m0-4.731v18.924M7.5 22.27a4.73 4.73 0 0 0 4.73-4.732"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}