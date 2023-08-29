package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5a.5.5 0 0 1 .5-.5h5.793L12 4.707V12.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-10ZM3.5 1A1.5 1.5 0 0 0 2 2.5v10A1.5 1.5 0 0 0 3.5 14h8a1.5 1.5 0 0 0 1.5-1.5V4.604a.75.75 0 0 0-.22-.53L9.854 1.145A.5.5 0 0 0 9.5 1h-6Zm1.75 6a.5.5 0 0 0 0 1h4.5a.5.5 0 0 0 0-1h-4.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}