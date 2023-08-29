package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M160 85q-18 0-30.5-12.5t-12.5-30t12.5-30T160 0t30.5 12.5t12.5 30t-12.5 30T160 85zm-79 73L21 459h45l39-171l44 43v128h43V299l-45-43l13-64q46 53 117 53v-42q-29 0-53.5-14.5T186 151l-22-34q-14-21-36-21q-3 0-8.5 1t-8.5 1L0 145v100h43v-72l38-15z"/>`),
		g.Group(children),
	)
}