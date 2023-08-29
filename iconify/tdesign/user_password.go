package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserPassword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0ZM17 12h4v2h-2v2.003h1.953L22.5 17.14v5.36h-9v-5.36l1.547-1.137H17V12Zm-9 4a4 4 0 0 0-4 4h7.55v2H2v-2a6 6 0 0 1 6-6h3.5v2H8Zm7.5 2.152V20.5h5v-2.348l-.203-.15h-4.594l-.203.15Z"/>`),
		g.Group(children),
	)
}