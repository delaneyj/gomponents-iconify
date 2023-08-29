package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyMobifone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 38.32V20.65m-4.875 14.624V20.649M43.5 38.32l-4.875-3.046M4.5 38.32l4.875-3.046m17.062 3.046l-4.875-3.046M4.5 38.32V20.65m4.875 14.624V20.649M26.437 38.32V20.65m-4.875 14.624V20.649m-12.187 0a6.094 6.094 0 1 1 12.187 0m4.875 0a6.094 6.094 0 1 1 12.188 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 20.649a10.969 10.969 0 0 1 19.499-6.896m.002 0A10.969 10.969 0 0 1 43.5 20.648"/>`),
		g.Group(children),
	)
}