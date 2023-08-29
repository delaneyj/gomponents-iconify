package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetrinaAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M6 8a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v2H6V8Zm6 20v-3H6V12h12v11h24c0 8.43-5.703 13.35-13 13.94V37h-5.5v2H29v3H14.5V30c-.747 0-2.5-.5-2.5-2Z"/><path d="M20 20a1 1 0 0 1 1-1h20a1 1 0 1 1 0 2H21a1 1 0 0 1-1-1Z"/></g>`),
		g.Group(children),
	)
}