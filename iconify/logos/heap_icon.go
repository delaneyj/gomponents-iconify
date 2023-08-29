package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeapIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M0 89.613h47.999v235.196H0zM104.025 0h47.999v179.196h-47.999z"/><path fill="#31D891" d="M104.025 235.229h47.999v179.196h-47.999zM208.001 89.613H256v235.196h-47.999z"/>`),
		g.Group(children),
	)
}