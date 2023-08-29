package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.026 9.965a5 5 0 0 0-6.992-6.992l6.992 6.992Zm-1.061 1.061a5 5 0 0 1-6.992-6.992l6.992 6.992Zm10.979 5.224a5 5 0 0 0-9.887 0h9.887Zm0 1.5h-9.888a5 5 0 0 0 9.887 0Z"/>`),
		g.Group(children),
	)
}