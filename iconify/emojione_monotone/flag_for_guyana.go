package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGuyana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m27.445 24.452c.304 1.501.479 3.046.531 4.622L15.334 9.528a28.172 28.172 0 0 1 3.949-2.461l40.162 19.385M14.808 52.441l19.026-19.027L35.249 32l-1.415-1.414l-19.026-19.027v-.063L57.293 32L14.808 52.505v-.064m4.475 4.493a28.179 28.179 0 0 1-3.949-2.462l44.642-21.546a28.02 28.02 0 0 1-.531 4.622L19.283 56.934"/>`),
		g.Group(children),
	)
}