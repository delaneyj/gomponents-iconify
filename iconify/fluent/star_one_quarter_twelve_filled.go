package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOneQuarterTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m5 2.12l-.828 1.678l-2.486.362a.798.798 0 0 0-.444 1.364l1.8 1.754l-.425 2.475a.837.837 0 0 0-.005.254a.801.801 0 0 0 1.165.59L5 9.954V2.12Z"/>`),
		g.Group(children),
	)
}