package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitchLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.301 3h16.7v11.7l-4.7 4.7h-3.9l-2.5 2.4h-2.9v-2.4h-4V6.2l1.3-3.2Zm.7 14.4h4v2.4h.095l2.5-2.4h3.877L19 13.872V5H5v12.4Zm10-9.4h2v4.7h-2V8Zm0 0h2v4.7h-2V8Zm-5 0h2v4.7h-2V8Z"/>`),
		g.Group(children),
	)
}