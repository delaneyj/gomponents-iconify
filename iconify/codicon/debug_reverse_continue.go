package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugReverseContinue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12v12h1.5V2zm-4.936.39L9.75 3v10l-1.186.61l-7-5V7.39l7-5zM3.29 8l4.96 3.543V4.457L3.29 8z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}