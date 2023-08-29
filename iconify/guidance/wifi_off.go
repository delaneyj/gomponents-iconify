package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5.5l23 23m-7.305-7.305l2.265-3.16A9.466 9.466 0 0 0 12 10.5a9.59 9.59 0 0 0-1.398.102m-2.683.817a9.517 9.517 0 0 0-2.379 1.615L8.381 17h.265A4.489 4.489 0 0 1 12 15.5M8.115 8.115A12.49 12.49 0 0 1 12 7.5c3.146 0 6.02 1.162 8.218 3.08L23 6.5v-.25C20.09 4.051 16.5 2.5 12 2.5c-2.954 0-5.516.668-7.75 1.75M2.022 5.522c-.35.233-.69.476-1.023.727v.25l2.782 4.08A12.6 12.6 0 0 1 5.7 9.199M10.5 20a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`),
		g.Group(children),
	)
}