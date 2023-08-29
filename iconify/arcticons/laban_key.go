package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabanKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.416 39.87v4.618c16.52.367 30.7-8.077 30.7-22.514c0-8.077-6.372-18.474-17.958-18.474c-13.893 0-20.273 9.169-20.273 18.474c0 13.848 16.475 16.17 19.115 6.35h-8.689V15.047h17.957c0 26.555-15.635 24.824-20.852 24.824v0Z"/>`),
		g.Group(children),
	)
}