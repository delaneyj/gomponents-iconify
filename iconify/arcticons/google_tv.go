package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.78 39.5H9.22a4.569 4.569 0 0 1-4.595-4.532V13.032A4.569 4.569 0 0 1 9.22 8.5h29.56a4.569 4.569 0 0 1 4.595 4.532v21.936A4.569 4.569 0 0 1 38.78 39.5Zm-28.343-5.813h27.126V14.314H10.436Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.375 17.219a2.906 2.906 0 0 0-2.906-2.906h-3.291M34.656 39.5a2.906 2.906 0 0 0 2.906-2.906v-3.291M4.625 30.781a2.906 2.906 0 0 0 2.906 2.906h3.291M13.344 8.5a2.906 2.906 0 0 0-2.906 2.906v3.291"/>`),
		g.Group(children),
	)
}