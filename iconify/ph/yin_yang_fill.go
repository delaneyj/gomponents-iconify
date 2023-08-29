package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYangFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 80a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm92 48A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-92 48a12 12 0 1 0-12 12a12 12 0 0 0 12-12Zm32-92a44.05 44.05 0 0 0-44-44a88 88 0 0 0-46.91 162.42A52 52 0 0 1 128 128a44.05 44.05 0 0 0 44-44Z"/>`),
		g.Group(children),
	)
}