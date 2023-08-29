package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 6.308l1.176 3.167l.347.936l.997.041l3.374.139l-2.647 2.092l-.784.62l.27.962l.911 3.249l-2.814-1.871l-.83-.553l-.83.552l-2.814 1.871l.911-3.249l.27-.962l-.784-.62l-2.648-2.092l3.374-.139l.997-.041l.347-.936L12 6.308M12 2L9.418 8.953L2 9.257l5.822 4.602L5.82 21L12 16.891L18.18 21l-2.002-7.141L22 9.257l-7.418-.305L12 2z"/>`),
		g.Group(children),
	)
}