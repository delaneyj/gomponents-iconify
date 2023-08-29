package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blinkist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M31.442 23.784a7.442 7.442 0 0 1-14.873 0C16.57 19.673 24 13.698 24 13.698s7.442 5.975 7.442 10.086Z"/><path d="M45.5 23.784c-.003 11.874-9.632 21.497-21.506 21.494C12.12 45.275 2.497 35.647 2.5 23.773a21.5 21.5 0 0 1 10.587-18.52a2.33 2.33 0 0 1 2.83.362l8.036 8.071l8.153-8.071a2.33 2.33 0 0 1 2.83-.361A21.512 21.512 0 0 1 45.5 23.784Z"/></g>`),
		g.Group(children),
	)
}