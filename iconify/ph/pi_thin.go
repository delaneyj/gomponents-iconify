package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PiThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 172a32 32 0 0 1-64 0V68H92v132a4 4 0 0 1-8 0V68H72a44.05 44.05 0 0 0-44 44a4 4 0 0 1-8 0a52.06 52.06 0 0 1 52-52h152a4 4 0 0 1 0 8h-52v104a24 24 0 0 0 48 0a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}