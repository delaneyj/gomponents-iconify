package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm4 191.91V148h28a4 4 0 0 0 0-8h-28v-28a20 20 0 0 1 20-20h16a4 4 0 0 0 0-8h-16a28 28 0 0 0-28 28v28H96a4 4 0 0 0 0 8h28v71.91a92 92 0 1 1 8 0Z"/>`),
		g.Group(children),
	)
}