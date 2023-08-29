package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockersLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M190 72a6 6 0 0 1-6 6h-24a6 6 0 0 1 0-12h24a6 6 0 0 1 6 6Zm-6 26h-24a6 6 0 0 0 0 12h24a6 6 0 0 0 0-12ZM72 78h24a6 6 0 0 0 0-12H72a6 6 0 0 0 0 12Zm24 20H72a6 6 0 0 0 0 12h24a6 6 0 0 0 0-12Zm126-50v176a6 6 0 0 1-12 0v-18h-76v18a6 6 0 0 1-12 0v-18H46v18a6 6 0 0 1-12 0V48a14 14 0 0 1 14-14h160a14 14 0 0 1 14 14ZM122 194V46H48a2 2 0 0 0-2 2v146Zm12 0h76V48a2 2 0 0 0-2-2h-74Z"/>`),
		g.Group(children),
	)
}