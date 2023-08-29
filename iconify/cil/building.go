package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M440 464V16H72v448H16v32h480v-32Zm-32 0H272v-64h-32v64H104V48h304Z"/><path fill="currentColor" d="M160 304h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32zm-160-96h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32zm-160-96h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32z"/>`),
		g.Group(children),
	)
}