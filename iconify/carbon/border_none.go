package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 3h2v2H3zm4 0h2v2H7zm4 0h2v2h-2zm4 0h2v2h-2zM3 7h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3zM27 3h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zM7 27h2v2H7zm4 0h2v2h-2zm4 0h2v2h-2zm8 0h2v2h-2zm-4 0h2v2h-2zm8-20h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zM8 10h10v2H8zm0 5h6v2H8z"/>`),
		g.Group(children),
	)
}