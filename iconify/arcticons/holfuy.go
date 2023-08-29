package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Holfuy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.25 34.79a19.5 19.5 0 1 1 30.634 3.713"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.052 42.69l-15.98-27.37L8.26 42.56l15.991-8.17l7.79 4.112"/>`),
		g.Group(children),
	)
}