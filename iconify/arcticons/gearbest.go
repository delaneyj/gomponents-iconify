package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gearbest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 14.972C7.573 26.442 15.141 33.028 24 33.028s16.425-6.58 19.5-18.056H33.094"/>`),
		g.Group(children),
	)
}