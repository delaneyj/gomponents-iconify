package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KoboBooks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><rect width="5.435" height="7.202" x="33.065" y="20.341" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.717"/><rect width="5.435" height="7.202" x="16.922" y="20.341" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.717"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 16.673v10.87m.001-2.309l4.92-4.897m-3.354 3.338l3.868 3.851m10.015-4.467a2.718 2.718 0 0 1 2.717-2.718h0a2.718 2.718 0 0 1 2.718 2.718v1.766a2.718 2.718 0 0 1-2.718 2.718h0a2.717 2.717 0 0 1-2.717-2.718m0 2.718v-10.87"/>`),
		g.Group(children),
	)
}