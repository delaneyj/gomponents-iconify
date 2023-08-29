package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockCountdownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 96a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm-12-24a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm28.66 56a8 8 0 0 0-8.63 7.31A88.12 88.12 0 1 1 120.66 40a8 8 0 0 0-1.32-16A104.12 104.12 0 1 0 232 136.66a8 8 0 0 0-7.34-8.66ZM128 56a72 72 0 1 1-72 72a72.08 72.08 0 0 1 72-72Zm-8 72a8 8 0 0 0 8 8h48a8 8 0 0 0 0-16h-40V80a8 8 0 0 0-16 0Zm40-80a12 12 0 1 0-12-12a12 12 0 0 0 12 12Z"/>`),
		g.Group(children),
	)
}