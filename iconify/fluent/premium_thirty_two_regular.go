package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PremiumThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 4a1 1 0 0 0-.894.553l-4 8a1 1 0 0 0 .118 1.078l13 16a1 1 0 0 0 1.552 0l13-16a1 1 0 0 0 .118-1.078l-4-8A1 1 0 0 0 25 4H7Zm-2.382 8l3-6h4.101l-1.5 6H4.618Zm.483 2h5.164l3.227 10.328L5.101 14Zm7.26 0h7.279L16 25.647L12.36 14Zm9.374 0H26.9l-8.391 10.328L21.735 14Zm5.647-2h-5.601l-1.5-6h4.101l3 6Zm-7.663 0h-7.438l1.5-6h4.438l1.5 6Z"/>`),
		g.Group(children),
	)
}