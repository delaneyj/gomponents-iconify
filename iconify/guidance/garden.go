package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Garden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 6.5V7m0-.5a6 6 0 0 1 6-6h.5V7a6.5 6.5 0 0 1-6.5 6.5m0-7a6 6 0 0 0-6-6h-.5V7a6.5 6.5 0 0 0 6.5 6.5m0 0V23m0-9.5v9m0 1H7.5a6 6 0 0 1-6-6H6a6 6 0 0 1 6 6Zm0 0h4.5a6 6 0 0 0 6-6H18a6 6 0 0 0-6 6Z"/>`),
		g.Group(children),
	)
}