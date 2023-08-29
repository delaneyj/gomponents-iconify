package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookLite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.5V18.575a5.075 5.075 0 0 1 5.075-5.075h0c2.498 0 4.057.74 5.126 2.123m-14.5 7.665h10.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 32.755V7.478A1.974 1.974 0 0 0 40.53 5.5H7.476A1.974 1.974 0 0 0 5.5 7.471v33.05A1.974 1.974 0 0 0 7.47 42.5h25.285"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.5 34.5v8h4"/>`),
		g.Group(children),
	)
}