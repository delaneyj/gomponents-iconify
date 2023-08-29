package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.955 6.7C5.92-1.301-2.006 39.293 13.817 34.03l19.304-7.132"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.948 21.068l18.448-6.412c17-6.038 8.89 34.538-17.557 26.676"/>`),
		g.Group(children),
	)
}