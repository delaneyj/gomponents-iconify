package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserIdentification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 11h4a1 1 0 0 1 1 1v2h-6v-2a1 1 0 0 1 1-1Z"/><circle cx="24" cy="8" r="2" fill="currentColor"/><path fill="currentColor" d="M30 18H18a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h12a2.002 2.002 0 0 1 2 2v12a2.003 2.003 0 0 1-2 2zM18 4v12h12.001L30 4zm-3 26h-2v-4a2.946 2.946 0 0 0-3-3H6a2.946 2.946 0 0 0-3 3v4H1v-4a4.951 4.951 0 0 1 5-5h4a4.951 4.951 0 0 1 5 5zM8 11a3 3 0 0 1 0 6a3 3 0 0 1 0-6m0-2a5 5 0 0 0 0 10A5 5 0 0 0 8 9z"/>`),
		g.Group(children),
	)
}