package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M80 384v32h216V78.627l84.687 84.687l22.626-22.628L280 17.373L156.687 140.686l22.626 22.628L264 78.627V384H80z"/>`),
		g.Group(children),
	)
}