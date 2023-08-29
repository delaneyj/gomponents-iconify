package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplaceFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M8 2H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm12 12h-4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2zM16.707 2.293a1 1 0 0 1 .083 1.32l-.083.094L15.414 5H19a3 3 0 0 1 2.995 2.824L22 8v3a1 1 0 0 1-1.993.117L20 11V8a1 1 0 0 0-.883-.993L19 7h-3.585l1.292 1.293a1 1 0 0 1-1.32 1.497l-.094-.083l-3-3a.98.98 0 0 1-.28-.872l.036-.146l.04-.104c.058-.126.14-.24.245-.334l2.959-2.958a1 1 0 0 1 1.414 0zM3 12a1 1 0 0 1 .993.883L4 13v3a1 1 0 0 0 .883.993L5 17h3.585l-1.292-1.293a1 1 0 0 1-.083-1.32l.083-.094a1 1 0 0 1 1.32-.083l.094.083l3 3a.98.98 0 0 1 .28.872l-.036.146l-.04.104a1.02 1.02 0 0 1-.245.334l-2.959 2.958a1 1 0 0 1-1.497-1.32l.083-.094L8.584 19H5a3 3 0 0 1-2.995-2.824L2 16v-3a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}