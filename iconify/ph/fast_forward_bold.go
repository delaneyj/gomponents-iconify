package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m242.81 111.29l-88.18-56.17A19.91 19.91 0 0 0 124 71.84v30L50.63 55.12A19.91 19.91 0 0 0 20 71.84v112.32a19.93 19.93 0 0 0 30.63 16.72L124 154.15v30a19.93 19.93 0 0 0 30.63 16.72l88.18-56.17a19.79 19.79 0 0 0 0-33.42ZM44 176.64V79.36L120.38 128Zm104 0V79.36L224.38 128Z"/>`),
		g.Group(children),
	)
}