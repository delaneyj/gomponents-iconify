package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lucia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="29.826" cy="18.167" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.826 21.584v7.454M8.56 17.744v9.882a1.334 1.334 0 0 0 1.412 1.411h.424m16.594-1.411a2.733 2.733 0 0 1-2.4 1.412h0a2.832 2.832 0 0 1-2.824-2.824V24.38a2.832 2.832 0 0 1 2.823-2.824h0a2.733 2.733 0 0 1 2.4 1.412m-8.649 3.246a2.832 2.832 0 0 1-2.824 2.824h0a2.832 2.832 0 0 1-2.824-2.824v-4.659"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.469 29.038a1.214 1.214 0 0 1-1.13-1.13v-6.353m19.971 4.659a2.832 2.832 0 0 1-2.823 2.824h0a2.832 2.832 0 0 1-2.824-2.824V24.38a2.832 2.832 0 0 1 2.824-2.823h0a2.832 2.832 0 0 1 2.823 2.823"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.44 29.038a1.214 1.214 0 0 1-1.13-1.13v-6.353"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}