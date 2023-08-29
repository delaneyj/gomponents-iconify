package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandSanitizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10V11a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3v10zm8-18H9a2 2 0 0 0-2 2m5-2v5m0 3v4m-2-2h4"/>`),
		g.Group(children),
	)
}