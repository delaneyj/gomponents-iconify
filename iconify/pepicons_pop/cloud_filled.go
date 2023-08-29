package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><rect width="16" height="8" x="2" y="6" rx="4"/><rect width="9" height="8" x="6" y="3" rx="4"/><rect width="16" height="8" x="2" y="6" rx="4"/><rect width="9" height="8" x="6" y="3" rx="4"/></g>`),
		g.Group(children),
	)
}