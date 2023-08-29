package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Netflix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29 25.75V5l8.4-.5v39L29 43M19 22.25V43l-8.4.5v-39h0L19 5m-.02 0l18.44 38.5m-26.84-39L29.02 43"/>`),
		g.Group(children),
	)
}