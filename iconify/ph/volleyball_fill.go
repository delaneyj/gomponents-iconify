package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolleyballFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm87.63 96h-34.26a104.18 104.18 0 0 0-35.78-78.23A88.18 88.18 0 0 1 215.63 120Zm-171.1 35.87a87.95 87.95 0 0 1 32.74-99.74l17.12 29.65a104.14 104.14 0 0 0-49.86 70.09Zm14.37 26.56a88 88 0 0 1 43.49-82.79L118.76 128l-41.49 71.87a88.62 88.62 0 0 1-18.37-17.44Zm150.84-21.85a88 88 0 0 1-93.49 3.78L132.62 136h83a87.16 87.16 0 0 1-5.88 24.58Z"/>`),
		g.Group(children),
	)
}