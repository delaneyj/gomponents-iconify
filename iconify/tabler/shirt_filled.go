package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShirtFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M14.883 3.007L14.978 3l.112.004l.113.017l.113.03l6 2a1 1 0 0 1 .677.833L22 6v5a1 1 0 0 1-.883.993L21 12h-2v7a2 2 0 0 1-1.85 1.995L17 21H7a2 2 0 0 1-1.995-1.85L5 19v-7H3a1 1 0 0 1-.993-.883L2 11V6a1 1 0 0 1 .576-.906l.108-.043l6-2A1 1 0 0 1 10 4a2 2 0 0 0 3.995.15l.009-.24l.017-.113l.037-.134l.044-.103l.05-.092l.068-.093l.069-.08c.056-.054.113-.1.175-.14l.096-.053l.103-.044l.108-.032l.112-.02z"/></g>`),
		g.Group(children),
	)
}