package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 48v48a4 4 0 0 1-8 0V57.66l-49.17 49.17a4 4 0 0 1-5.66-5.66L198.34 52H160a4 4 0 0 1 0-8h48a4 4 0 0 1 4 4ZM101.17 149.17L52 198.34V160a4 4 0 0 0-8 0v48a4 4 0 0 0 4 4h48a4 4 0 0 0 0-8H57.66l49.17-49.17a4 4 0 0 0-5.66-5.66ZM208 156a4 4 0 0 0-4 4v38.34l-49.17-49.17a4 4 0 0 0-5.66 5.66L198.34 204H160a4 4 0 0 0 0 8h48a4 4 0 0 0 4-4v-48a4 4 0 0 0-4-4ZM57.66 52H96a4 4 0 0 0 0-8H48a4 4 0 0 0-4 4v48a4 4 0 0 0 8 0V57.66l49.17 49.17a4 4 0 0 0 5.66-5.66Z"/>`),
		g.Group(children),
	)
}