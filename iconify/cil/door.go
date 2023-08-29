package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M400 464V64H112v400H16v32h480v-32Zm-32 0H144V96h224Z"/><path fill="currentColor" d="M312 252h32v72h-32z"/>`),
		g.Group(children),
	)
}