package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m223.57 81.81l-41.38-41.38a22 22 0 0 0-31.12 0L32.43 159.07a22 22 0 0 0 0 31.11l30.07 30.06a6 6 0 0 0 4.24 1.76H216a6 6 0 0 0 0-12h-89.51l97.08-97.08a22 22 0 0 0 0-31.11ZM109.51 210H69.22l-28.3-28.3a10 10 0 0 1 0-14.15L96 112.48L151.52 168Zm105.57-105.56L160 159.51L104.48 104l55.08-55.07a10 10 0 0 1 14.14 0l41.38 41.37a10 10 0 0 1 0 14.14Z"/>`),
		g.Group(children),
	)
}