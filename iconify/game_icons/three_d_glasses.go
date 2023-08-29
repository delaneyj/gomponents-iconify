package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDGlasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M176 80L39.36 247h77.74L176 144l32 48v-48l-32-64zm160 0l-32 64v48l32-48l58.9 103h77.7L336 80zM25 265v174h194.2l36.8-55.2l36.8 55.2H487V265H25zm23 23h176v64l-32 64H48V288zm240 0h176v128H320l-32-64v-64z"/>`),
		g.Group(children),
	)
}