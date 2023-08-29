package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 16c0 3.7-.6 6.7-1.4 6.7c-.8 0-1.4-3-1.4-6.7s.6-6.7 1.4-6.7c.8 0 1.4 3 1.4 6.7m-3.5 0c0 4.1-1.8 7.5-3.9 7.5s-3.9-3.4-3.9-7.5s1.8-7.5 3.9-7.5s3.9 3.4 3.9 7.5m-8.7 0c0 4.4-3.5 8-7.9 8S2 20.4 2 16s3.5-8 7.9-8s7.9 3.6 7.9 8"/>`),
		g.Group(children),
	)
}