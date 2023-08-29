package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchThemes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M23 4v23h20.993L44 4H23Z" clip-rule="evenodd"/><path d="M31.005 44H18.658c-1.702 0-3.742-.568-5.11-2.387c-.925-1.23-1.543-3.03-1.543-5.613V25"/><path stroke-linejoin="round" d="m4 33l8.005-8l8.009 8"/><path d="M30 19h7m-7-7h7"/></g>`),
		g.Group(children),
	)
}