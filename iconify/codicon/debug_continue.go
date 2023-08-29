package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugContinue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 2H4v12H2.5V2zm4.936.39L6.25 3v10l1.186.61l7-5V7.39l-7-5zM12.71 8l-4.96 3.543V4.457L12.71 8z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}