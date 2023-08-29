package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AfricanTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22V12m0-4v4m0 0l3-3m-2.576 9.576l6.169-6.169a5.502 5.502 0 0 0-.513-8.234a9.904 9.904 0 0 0-12.16 0a5.502 5.502 0 0 0-.513 8.234l6.169 6.169a.6.6 0 0 0 .848 0Z"/>`),
		g.Group(children),
	)
}