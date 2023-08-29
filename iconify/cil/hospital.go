package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M160 304h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32z"/><path fill="currentColor" d="M440 464V88h-72V16H144v72H72v376H16v32h480v-32ZM176 48h160v160H176Zm232 416H272v-64h-32v64H104V120h40v120h224V120h40Z"/><path fill="currentColor" d="M272 80h-32v32h-32v32h32v32h32v-32h32v-32h-32V80z"/>`),
		g.Group(children),
	)
}