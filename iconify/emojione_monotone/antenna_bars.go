package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntennaBars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2M20 48h-6v-8h6v8m10 0h-6V32h6v16m10 0h-6V24h6v24m10 0h-6V16h6v32"/>`),
		g.Group(children),
	)
}