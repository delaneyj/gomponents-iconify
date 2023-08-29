package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104 104 0 0 0 128 24Zm0 184a80 80 0 1 1 53.34-139.63a8 8 0 0 1-10.67 11.92A64 64 0 1 0 191.5 136H128a8 8 0 0 1 0-16h72a8 8 0 0 1 8 8a80.09 80.09 0 0 1-80 80Z"/>`),
		g.Group(children),
	)
}