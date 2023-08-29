package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upgradeall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.187 41.878q.471-.325.923-.673m-6.479 3.409a21.578 21.578 0 0 0 2.755-1.106M22.19 45.592a21.34 21.34 0 0 0 5.254-.205m-15.949-3.739a21.583 21.583 0 0 0 7.438 3.413M41.88 12.284a21.479 21.479 0 1 0-32.733 27.42"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.33 6.914l-.451 5.37l-5.371-.451"/>`),
		g.Group(children),
	)
}