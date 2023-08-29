package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstanceClassic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.5 21H23v-2h.5a4.497 4.497 0 0 0 .356-8.981l-.815-.064l-.099-.812a6.994 6.994 0 0 0-13.883 0l-.1.812l-.815.064A4.497 4.497 0 0 0 8.5 19H9v2h-.5A6.497 6.497 0 0 1 7.2 8.136a8.994 8.994 0 0 1 17.6 0A6.497 6.497 0 0 1 23.5 21Z"/><circle cx="9" cy="27" r="1" fill="currentColor"/><path fill="currentColor" d="M26 23h-9v-7.17l2.59 2.58L21 17l-5-5l-5 5l1.41 1.41L15 15.83V23H6a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h20a2.002 2.002 0 0 0 2-2v-4a2.002 2.002 0 0 0-2-2Zm0 6H6v-4h20Z"/>`),
		g.Group(children),
	)
}