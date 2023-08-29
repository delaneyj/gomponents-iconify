package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bewakoof(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="13.939" cy="26.02" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.439" ry="7.249"/><ellipse cx="34.061" cy="26.02" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.439" ry="7.249"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.017 15.187c4.832-1.175 6.693-.053 9.957 3.106"/>`),
		g.Group(children),
	)
}