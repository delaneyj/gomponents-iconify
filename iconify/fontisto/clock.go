package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12C23.98 5.381 18.619.02 12.002 0zm0 21.6c-5.302 0-9.6-4.298-9.6-9.6S6.698 2.4 12 2.4s9.6 4.298 9.6 9.6c-.016 5.296-4.304 9.584-9.598 9.6zM12.6 6h-1.8v7.2l6.24 3.84l.96-1.56l-5.4-3.24z"/>`),
		g.Group(children),
	)
}