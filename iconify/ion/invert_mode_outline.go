package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertModeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="256" cy="256" r="208" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"/><path fill="currentColor" d="M256 176v160a80 80 0 0 1 0-160Zm0-128v128a80 80 0 0 1 0 160v128c114.88 0 208-93.12 208-208S370.88 48 256 48Z"/>`),
		g.Group(children),
	)
}