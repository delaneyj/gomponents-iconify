package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PricetagsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M288 16L0 304l176 176l288-288V16Zm80 128a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/><path fill="currentColor" d="M480 64v144L216.9 471.1L242 496l270-272V64h-32z"/>`),
		g.Group(children),
	)
}