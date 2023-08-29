package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassOfMilk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7.79 29.193L4.01 3.066A.942.942 0 0 1 4.946 2h22.106a.942.942 0 0 1 .938 1.066l-3.78 26.127a.934.934 0 0 1-.938.807H8.727a.95.95 0 0 1-.937-.807Zm1.181-2.73c.04.308.31.537.62.537h10.122l2.723-19H7.64c-.69 0-1.22.617-1.13 1.295l2.461 17.167ZM21.733 27h.664c.31 0 .58-.229.62-.538l2.47-17.167a1.13 1.13 0 0 0-1.03-1.29l.399-2.79a1 1 0 1 0-1.98-.284L22.436 8h1.921a1.1 1.1 0 0 1 .1.004L21.732 27Z"/>`),
		g.Group(children),
	)
}