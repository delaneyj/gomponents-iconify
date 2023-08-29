package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 10c4.418 0 8-1.79 8-4s-3.582-4-8-4s-8 1.79-8 4s3.582 4 8 4Zm6.328.17A7.61 7.61 0 0 0 20 9.053V18c0 2.21-3.582 4-8 4s-8-1.79-8-4V9.053a7.61 7.61 0 0 0 1.672 1.117C7.37 11.018 9.608 11.5 12 11.5c2.392 0 4.63-.482 6.328-1.33Z"/>`),
		g.Group(children),
	)
}