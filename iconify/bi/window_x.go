package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M2.5 5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM4 5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/><path d="M0 4a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v4a.5.5 0 0 1-1 0V7H1v5a1 1 0 0 0 1 1h5.5a.5.5 0 0 1 0 1H2a2 2 0 0 1-2-2V4Zm1 2h13V4a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v2Z"/><path d="M16 12.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm-4.854-1.354a.5.5 0 0 0 0 .708l.647.646l-.647.646a.5.5 0 0 0 .708.708l.646-.647l.646.647a.5.5 0 0 0 .708-.708l-.647-.646l.647-.646a.5.5 0 0 0-.708-.708l-.646.647l-.646-.647a.5.5 0 0 0-.708 0Z"/></g>`),
		g.Group(children),
	)
}