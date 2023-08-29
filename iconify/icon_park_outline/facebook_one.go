package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" fill-rule="evenodd" clip-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4.111 46A2.117 2.117 0 0 1 2 43.877V4.123C2 2.95 2.945 2 4.111 2H43.89C45.055 2 46 2.95 46 4.123v39.754A2.117 2.117 0 0 1 43.889 46H4.11Z"/><path fill="currentColor" d="M38.662 17.887h-4.869a2.445 2.445 0 0 0-2.444 2.444v4.89h7.313l-1.071 7.332h-6.242V46h-8.042V32.553h-6.64v-7.331h6.554l.086-6.959l-.024-2.494a6.377 6.377 0 0 1 6.376-6.438h9.003v8.556Z"/></g>`),
		g.Group(children),
	)
}