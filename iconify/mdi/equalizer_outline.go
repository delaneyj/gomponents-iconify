package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualizerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 21H9V3h6v18m-4-2h2V5h-2v14m-3 2H2V11h6v10m-4-2h2v-6H4v6m18 2h-6V8h6v13m-4-2h2v-9h-2v9Z"/>`),
		g.Group(children),
	)
}