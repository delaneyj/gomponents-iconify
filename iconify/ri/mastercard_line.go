package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MastercardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 18.294a7.3 7.3 0 1 1 0-12.588a7.3 7.3 0 1 1 0 12.588Zm1.702-1.383a5.3 5.3 0 1 0 0-9.82A7.273 7.273 0 0 1 15.6 12c0 1.89-.719 3.613-1.898 4.91ZM10.299 7.09a5.3 5.3 0 1 0 0 9.82A7.274 7.274 0 0 1 8.401 12c0-1.89.719-3.614 1.898-4.91Zm1.702 1.115a5.284 5.284 0 0 0-1.6 3.795c0 1.488.613 2.832 1.6 3.795a5.284 5.284 0 0 0 1.6-3.795a5.284 5.284 0 0 0-1.6-3.795Z"/>`),
		g.Group(children),
	)
}