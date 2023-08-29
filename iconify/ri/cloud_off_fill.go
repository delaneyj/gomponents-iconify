package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.515 2.1l19.092 19.092l-1.415 1.415l-2.014-2.015A5.984 5.984 0 0 1 17 21H7A6 6 0 0 1 5.008 9.339a6.992 6.992 0 0 1 .353-2.563L2.1 3.514L3.515 2.1ZM17 9a6.003 6.003 0 0 1 5.204 8.989L14.01 9.796C14.89 9.29 15.91 9 17 9Zm-5-7a7.003 7.003 0 0 1 6.765 5.195a8.027 8.027 0 0 0-6.206 1.15L7.694 3.48A6.97 6.97 0 0 1 12 2Z"/>`),
		g.Group(children),
	)
}