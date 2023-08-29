package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExternalLink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.21 1.5H10v1.25h3.08L7.9 7.21l.82 1l5.53-4.77V7h1.25V2.79a1.29 1.29 0 0 0-1.29-1.29z"/><path fill="currentColor" d="M12.25 13.25H1.75v-8.5H7.5V3.5h-6a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-4h-1.25z"/>`),
		g.Group(children),
	)
}