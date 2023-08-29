package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trackworktime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 20.43a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v17.3a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2ZM24 9.41a9 9 0 0 1 9 9H15a9 9 0 0 1 9-9Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.6 20.46a8.62 8.62 0 1 1-8.6 8.62a8.62 8.62 0 0 1 8.6-8.62Zm.1 8.62l-4.67 4.56m4.67-10.07v5.51"/>`),
		g.Group(children),
	)
}