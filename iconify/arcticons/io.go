package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Io(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="30.597" cy="26.455" r="12.903" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="6.561" height="20.047" x="4.5" y="19.311" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.28"/><circle cx="7.78" cy="11.922" r="3.28" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}