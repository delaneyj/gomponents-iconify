package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSettingsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M43 23v-9a2 2 0 0 0-2-2H24l-5-6H7a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h15"/><circle cx="35" cy="35" r="4"/><path stroke-linecap="round" stroke-linejoin="round" d="M35 28v3m0 8v3m4.828-12l-2.121 2.121M31.828 38l-2.121 2.121M30 30l2.121 2.121M38 38l2.121 2.121M28 35h3m8 0h3"/></g>`),
		g.Group(children),
	)
}