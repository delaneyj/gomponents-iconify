package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSeychelles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2M17.148 55.72c-.241-.151-.482-.302-.718-.46l42.34-15.053a28.02 28.02 0 0 1-6.652 11.238l-34.97 4.275m39.291-37.362L14.975 54L36.121 4.306c8.734 1.294 16.162 6.636 20.318 14.052"/>`),
		g.Group(children),
	)
}