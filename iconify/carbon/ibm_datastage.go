package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmDatastage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 22.1V13h-2v9.1c-1.4.4-2.5 1.5-2.9 2.9H13v2h9.1c.4 1.7 2 3 3.9 3c2.2 0 4-1.8 4-4c0-1.9-1.3-3.4-3-3.9zM26 28c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2zm-16 2H2v-8h8v8zm-6-2h4v-4H4v4zM9.9 7H19V5H9.9c-.5-1.7-2-3-3.9-3c-2.2 0-4 1.8-4 4c0 1.9 1.3 3.4 3 3.9V19h2V9.9C8.4 9.5 9.5 8.4 9.9 7zM6 8c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2z"/><path fill="currentColor" d="M22 2v2h4.6L11 19.6l1.4 1.4L28 5.4V10h2V2z"/>`),
		g.Group(children),
	)
}