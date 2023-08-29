package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FetchUploadCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 13l-5 5l1.41 1.41L15 16.83V28H6v2h9a2 2 0 0 0 2-2V16.83l2.59 2.58L21 18Z"/><path fill="currentColor" d="M23.5 22H23v-2h.5a4.497 4.497 0 0 0 .356-8.981l-.815-.064l-.099-.812a6.994 6.994 0 0 0-13.883 0l-.1.812l-.815.064A4.497 4.497 0 0 0 8.5 20H9v2h-.5A6.497 6.497 0 0 1 7.2 9.136a8.994 8.994 0 0 1 17.6 0A6.497 6.497 0 0 1 23.5 22Z"/>`),
		g.Group(children),
	)
}