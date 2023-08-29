package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBangladesh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m-5.9 43.872c-7.662 0-13.874-6.211-13.874-13.872c0-7.663 6.211-13.874 13.874-13.874c7.664 0 13.873 6.211 13.873 13.874c0 7.661-6.209 13.872-13.873 13.872"/>`),
		g.Group(children),
	)
}