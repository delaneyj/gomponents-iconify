package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gbp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1020 1009v367q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-150q0-13 9.5-22.5T32 1194h97V811H34q-14 0-23-9.5T2 779V648q0-14 9-23t23-9h95V393q0-171 123.5-282T567 0q185 0 335 125q9 8 10 20.5t-7 22.5L802 295q-9 11-22 12q-13 2-23-7q-5-5-26-19t-69-32t-93-18q-85 0-137 47t-52 123v215h305q13 0 22.5 9t9.5 23v131q0 13-9.5 22.5T685 811H380v379h414v-181q0-13 9-22.5t23-9.5h162q14 0 23 9.5t9 22.5z"/>`),
		g.Group(children),
	)
}