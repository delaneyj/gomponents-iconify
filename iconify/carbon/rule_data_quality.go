package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleDataQuality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m30 28.6l-2.8-2.8c.5-.8.8-1.8.8-2.8c0-2.8-2.2-5-5-5s-5 2.2-5 5s2.2 5 5 5c1 0 2-.3 2.8-.8l2.8 2.8l1.4-1.4zM20 23c0-1.7 1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3s-3-1.3-3-3zM8 16h10v2H8zm0-6h12v2H8z"/><path fill="currentColor" d="m14 27.7l-5.2-2.8C5.8 23.4 4 20.3 4 17V4h20v11h2V4c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v13c0 4.1 2.2 7.8 5.8 9.7L14 30v-2.3z"/>`),
		g.Group(children),
	)
}