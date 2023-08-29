package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Packaging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="m11.656.656l-1.312.688l1 2l-2-1l-.688 1.312L11.344 5H2a1 1 0 0 0-.094 0a1 1 0 0 0-.093 0A1 1 0 0 0 1 6v4a1 1 0 0 0 1 1v13a1 1 0 0 0 1 1h20a1 1 0 0 0 1-1V11a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-9.344l2.688-1.344l-.688-1.312l-2 1l1-2l-1.312-.688L13 3.344L11.656.656zM3 7h9v2H3.187A1 1 0 0 0 3 8.969V7zm11 0h9v2h-9V7zM4 11h8v12H4V11zm10 0h8v12h-8V11z"/>`),
		g.Group(children),
	)
}