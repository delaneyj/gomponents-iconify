package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrademarkThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm-20-116a4 4 0 0 1-4 4H92v44a4 4 0 0 1-8 0v-44H72a4 4 0 0 1 0-8h32a4 4 0 0 1 4 4Zm80 0v48a4 4 0 0 1-8 0v-37.35l-21 24a4 4 0 0 1-6 0l-21-24V152a4 4 0 0 1-8 0v-48a4 4 0 0 1 7-2.63l25 28.56l25-28.56a4 4 0 0 1 7 2.63Z"/>`),
		g.Group(children),
	)
}