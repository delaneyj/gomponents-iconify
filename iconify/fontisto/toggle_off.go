package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 11.617A6.818 6.818 0 0 1 6.813 4.8h10.371a6.817 6.817 0 1 1 0 13.634H6.818a6.818 6.818 0 0 1-6.817-6.813zm6.817 4.543a4.543 4.543 0 1 0-.003-9.085a4.543 4.543 0 0 0 .001 9.085h.003z"/>`),
		g.Group(children),
	)
}