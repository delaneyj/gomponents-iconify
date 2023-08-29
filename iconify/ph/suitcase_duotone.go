package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M40 64h48v144H40a8 8 0 0 1-8-8V72a8 8 0 0 1 8-8Zm176 0h-48v144h48a8 8 0 0 0 8-8V72a8 8 0 0 0-8-8Z" opacity=".2"/><path d="M216 56h-40v-8a24 24 0 0 0-24-24h-48a24 24 0 0 0-24 24v8H40a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V72a16 16 0 0 0-16-16ZM96 48a8 8 0 0 1 8-8h48a8 8 0 0 1 8 8v8H96Zm64 24v128H96V72ZM40 72h40v128H40Zm176 128h-40V72h40v128Z"/></g>`),
		g.Group(children),
	)
}