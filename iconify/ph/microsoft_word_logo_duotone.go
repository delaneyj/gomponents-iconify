package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftWordLogoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M152 80v96a8 8 0 0 1-8 8H40a8 8 0 0 1-8-8V80a8 8 0 0 1 8-8h104a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M200 24H72a16 16 0 0 0-16 16v24H40a16 16 0 0 0-16 16v96a16 16 0 0 0 16 16h16v24a16 16 0 0 0 16 16h128a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16Zm-40 80h40v48h-40ZM72 40h128v48h-40v-8a16 16 0 0 0-16-16H72ZM40 80h104v96H40Zm32 136v-24h72a16 16 0 0 0 16-16v-8h40v48Zm-3.76-62.06l-12-48a8 8 0 1 1 15.52-3.88l6.76 27l6.32-12.66a8 8 0 0 1 14.32 0l6.32 12.66l6.76-27a8 8 0 0 1 15.52 3.88l-12 48a8 8 0 0 1-6.89 6a8.46 8.46 0 0 1-.87.05a8 8 0 0 1-7.16-4.42L92 137.89l-8.84 17.69a8 8 0 0 1-14.92-1.64Z"/></g>`),
		g.Group(children),
	)
}