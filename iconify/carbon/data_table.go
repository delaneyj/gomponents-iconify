package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 18h4v2H8zm6 0h4v2h-4zm-6-4h4v2H8zm6 8h4v2h-4zm6-8h4v2h-4zm0 8h4v2h-4z"/><path fill="currentColor" d="M27 3H5a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h22a2.002 2.002 0 0 0 2-2V5a2.002 2.002 0 0 0-2-2Zm0 2v4H5V5ZM5 27V11h22v16Z"/>`),
		g.Group(children),
	)
}