package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnowflakeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 2v16m0 4v-4m0 0l3 3m-3-3l-3 3m6-18l-3 3l-3-3M3.34 7l3.464 2m0 0l10.392 6M6.804 9L5.706 4.902M6.804 9l-4.098 1.098M17.196 15l3.464 2m-3.464-2l4.098-1.098M17.196 15l1.098 4.099M20.66 7l-3.464 2m0 0L6.804 15m10.392-6l1.098-4.098M17.196 9l4.098 1.098M6.804 15L3.34 17m3.464-2l-4.098-1.098M6.804 15l-1.098 4.1"/>`),
		g.Group(children),
	)
}