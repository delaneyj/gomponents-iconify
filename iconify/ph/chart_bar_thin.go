package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 204h-12V40a4 4 0 0 0-4-4h-56a4 4 0 0 0-4 4v44H96a4 4 0 0 0-4 4v44H48a4 4 0 0 0-4 4v68H32a4 4 0 0 0 0 8h192a4 4 0 0 0 0-8ZM156 44h48v160h-48Zm-56 48h48v112h-48Zm-48 48h40v64H52Z"/>`),
		g.Group(children),
	)
}