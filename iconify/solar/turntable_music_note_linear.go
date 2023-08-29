package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurntableMusicNoteLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M12 22c-4.714 0-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464c1.241 1.241 1.43 3.123 1.46 6.536M19 20v-8"/><circle cx="17" cy="20" r="2"/><path stroke-linecap="round" d="M22 15a3 3 0 0 1-3-3"/><path d="M9 12a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/></g>`),
		g.Group(children),
	)
}