package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarThreeQuarterTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m17 7.349l-1.698-3.441c-.532-1.078-2.069-1.078-2.6 0L10.01 9.36l-6.017.875c-1.19.173-1.664 1.634-.804 2.473l4.355 4.244l-1.028 5.993c-.203 1.185 1.04 2.088 2.104 1.529l5.382-2.83L17 23.221V7.349Z"/>`),
		g.Group(children),
	)
}