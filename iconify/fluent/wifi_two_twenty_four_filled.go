package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiTwoTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.96 16.44a1.501 1.501 0 1 1 2.122 2.122a1.501 1.501 0 0 1-2.123-2.122Zm-2.604-2.742a5.232 5.232 0 0 1 7.4 0c.46.461.838 1.025 1.101 1.625a1 1 0 1 1-1.832.803a3.356 3.356 0 0 0-.683-1.013a3.233 3.233 0 0 0-4.572 0a3.255 3.255 0 0 0-.672 1a1 1 0 1 1-1.832-.802a5.25 5.25 0 0 1 1.09-1.613ZM6.31 10.707a8.128 8.128 0 0 1 11.495 0a8.35 8.35 0 0 1 1.504 2.085a1 1 0 1 1-1.781.91a6.36 6.36 0 0 0-1.137-1.581a6.128 6.128 0 0 0-9.8 1.562a1 1 0 1 1-1.784-.902a8.08 8.08 0 0 1 1.503-2.074Z"/>`),
		g.Group(children),
	)
}