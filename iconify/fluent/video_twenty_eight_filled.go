package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 5.5A3.25 3.25 0 0 0 2 8.75v10.5a3.25 3.25 0 0 0 3.25 3.25h9.5A3.25 3.25 0 0 0 18 19.25V8.75a3.25 3.25 0 0 0-3.25-3.25h-9.5Zm17.873 15.143l-3.623-3.55V11l3.612-3.628c.787-.79 2.136-.233 2.136.882V19.75c0 1.108-1.334 1.668-2.125.893Z"/>`),
		g.Group(children),
	)
}