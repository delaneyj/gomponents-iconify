package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadiofmFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.28 38.417v-9.501l4.756 9.512l4.756-9.498v9.498m-16.694-9.512h4.756m-4.756 4.756h3.091m-3.091-4.756v9.512m24.206 0l-.001-9.512l-5.105 6.389H38.5"/><rect width="4" height="5.3" x="30.205" y="21.456" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.098 23.456a2 2 0 0 1 2-2h0m-2 0v5.3"/><circle cx="28.301" cy="18.931" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.301 21.381v5.3m-7.802-1.925a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2m0 3.3v-5.3m5.805 2a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2v1.3a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2m0 2v-8"/>`),
		g.Group(children),
	)
}