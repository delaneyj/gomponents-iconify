package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiThreeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.938 16.446a1.501 1.501 0 1 1 2.123 2.123a1.501 1.501 0 0 1-2.123-2.123Zm-2.603-2.742a5.233 5.233 0 0 1 7.4 0c.46.461.838 1.025 1.101 1.625a1 1 0 0 1-1.832.803a3.356 3.356 0 0 0-.683-1.013a3.233 3.233 0 0 0-4.572 0a3.255 3.255 0 0 0-.672 1a1 1 0 1 1-1.832-.802a5.25 5.25 0 0 1 1.09-1.613Z"/>`),
		g.Group(children),
	)
}