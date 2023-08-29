package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteTvOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5.27L3.28 4L20 20.72L18.73 22L17 20.25c-.14.99-1 1.75-2 1.75H9a2 2 0 0 1-2-2v-9.73l-5-5M9 2h2v2h2V2h2a2 2 0 0 1 2 2v11.18l-4-4V10h2V8h-2V6h-2v2H9.82L7 5.18V4a2 2 0 0 1 2-2m0 18h2v-2H9v2m4 0h2v-1.73l-.27-.27H13v2m-4-6v2h2v-1.73l-.27-.27H9Z"/>`),
		g.Group(children),
	)
}