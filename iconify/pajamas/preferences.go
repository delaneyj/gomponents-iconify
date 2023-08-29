package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Preferences(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM12 8a3 3 0 1 0-2.905-3.75H1.75a.75.75 0 0 0 0 1.5h7.345A3.001 3.001 0 0 0 12 8Zm-6.5 3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm1.405.75A3.001 3.001 0 0 1 1 11a3 3 0 0 1 5.905-.75h7.345a.75.75 0 0 1 0 1.5H6.905Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}