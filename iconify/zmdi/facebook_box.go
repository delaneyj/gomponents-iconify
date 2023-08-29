package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M363 88h-54q-31 0-52.5 22T235 163v53h-43v64h43v149h64V280h64v-64h-64v-43q0-8 6-14.5t15-6.5h43V88zM0 3h427v426H0V3z"/>`),
		g.Group(children),
	)
}