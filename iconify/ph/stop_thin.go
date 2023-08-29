package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200.73 44H55.27A11.29 11.29 0 0 0 44 55.27v145.46A11.29 11.29 0 0 0 55.27 212h145.46A11.29 11.29 0 0 0 212 200.73V55.27A11.29 11.29 0 0 0 200.73 44ZM204 200.73a3.28 3.28 0 0 1-3.27 3.27H55.27a3.28 3.28 0 0 1-3.27-3.27V55.27A3.28 3.28 0 0 1 55.27 52h145.46a3.28 3.28 0 0 1 3.27 3.27Z"/>`),
		g.Group(children),
	)
}