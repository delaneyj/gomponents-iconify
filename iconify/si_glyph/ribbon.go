package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ribbon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m9.025 0l1.39 1.025h1.717l.53 1.657l1.39 1.025l-.531 1.659l.531 1.658l-1.39 1.024l-.53 1.658h-1.717l-1.39 1.025l-1.39-1.025H5.92l-.529-1.658l-1.389-1.024l.529-1.658l-.529-1.659l1.389-1.025l.529-1.657h1.715L9.025 0zM5.042 15.529l-1.04-2.682l-1.984 1.599L4.006 9l.675 1.577l2.317.014l-1.956 4.938zm7.914-.021l1.075-2.661l1.948 1.578L14.027 9l-.675 1.577l-2.317.014l1.921 4.917z"/>`),
		g.Group(children),
	)
}