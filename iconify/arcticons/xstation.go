package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xstation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.56 9.48v28.8h7.09l10.06-14.91l-9.37-13.89Zm18.44 0L33.1 17l5.1-7.56Zm5.1 19.87l-6 8.93h12Z"/><rect width="37" height="37" x="5.72" y="5.69" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}