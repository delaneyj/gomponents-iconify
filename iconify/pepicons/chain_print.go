package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChainPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><g fill="currentColor" stroke-width="2" opacity=".8"><rect width="6" height="10" x="14.784" y="3.384" rx="3" transform="rotate(33.038 14.784 3.384)"/><rect width="6" height="10" x="9.836" y="7.323" rx="3" transform="rotate(33.038 9.836 7.323)"/></g><rect width="7" height="11" x="13.137" y="1.192" rx="3.5" transform="rotate(33.038 13.137 1.192)"/><rect width="7" height="11" x="8.189" y="5.131" rx="3.5" transform="rotate(33.038 8.189 5.131)"/></g>`),
		g.Group(children),
	)
}