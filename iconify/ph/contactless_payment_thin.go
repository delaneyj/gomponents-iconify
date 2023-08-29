package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactlessPaymentThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M93.54 102.13a55.39 55.39 0 0 1 0 51.74A4 4 0 0 1 90 156a4.07 4.07 0 0 1-1.87-.46a4 4 0 0 1-1.67-5.41a46.73 46.73 0 0 0 0-44.26a4 4 0 1 1 7.08-3.74Zm50.58-33.66a4 4 0 0 0-1.65 5.41a114.67 114.67 0 0 1 0 108.24a4 4 0 1 0 7.06 3.76a122.65 122.65 0 0 0 0-115.76a4 4 0 0 0-5.41-1.65Zm-28 16a4 4 0 0 0-1.65 5.41a81 81 0 0 1 0 76.24a4 4 0 1 0 7.06 3.76a89 89 0 0 0 0-83.76a4 4 0 0 0-5.41-1.65ZM228 128A100 100 0 1 1 128 28a100.11 100.11 0 0 1 100 100Zm-8 0a92 92 0 1 0-92 92a92.1 92.1 0 0 0 92-92Z"/>`),
		g.Group(children),
	)
}