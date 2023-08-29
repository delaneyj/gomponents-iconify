package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaveSawtoothLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m235.14 133.11l-104 64A6 6 0 0 1 122 192V74.74l-94.85 58.37a6 6 0 1 1-6.29-10.22l104-64A6 6 0 0 1 134 64v117.26l94.85-58.37a6 6 0 1 1 6.29 10.22Z"/>`),
		g.Group(children),
	)
}