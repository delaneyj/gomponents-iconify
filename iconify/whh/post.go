package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Post(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M855.048 768h-215l-256 256V768h-215q-57 0-113-57t-56-115V172q0-58 56-115t113-57h686q57 0 113 57t56 115v424q0 58-56 115t-113 57zm41-541q0-41-29-70t-69-29h-572q-41 0-69.5 29t-28.5 70v314q0 41 28.5 70t69.5 29h572q40 0 69-29t29-70V227z"/>`),
		g.Group(children),
	)
}