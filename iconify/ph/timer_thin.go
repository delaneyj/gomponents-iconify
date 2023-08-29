package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 44a92 92 0 1 0 92 92a92.1 92.1 0 0 0-92-92Zm0 176a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm42.83-126.83a4 4 0 0 1 0 5.66l-40 40a4 4 0 1 1-5.66-5.66l40-40a4 4 0 0 1 5.66 0ZM100 16a4 4 0 0 1 4-4h48a4 4 0 0 1 0 8h-48a4 4 0 0 1-4-4Z"/>`),
		g.Group(children),
	)
}