package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="21.5" cy="29.5" r="1.5" fill="currentColor"/><circle cx="24.5" cy="25.5" r="1.5" fill="currentColor"/><circle cx="7.5" cy="25.5" r="1.5" fill="currentColor"/><circle cx="4.5" cy="29.5" r="1.5" fill="currentColor"/><circle cx="10.499" cy="29.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="m15.868 30.496l-1.736-.992L17.276 24h-5.999l4.855-8.496l1.736.992L14.723 22h6.001l-4.856 8.496z"/><path fill="currentColor" d="M23.5 22H23v-2h.5a4.497 4.497 0 0 0 .356-8.981l-.815-.064l-.099-.812a6.994 6.994 0 0 0-13.883 0l-.1.812l-.815.064A4.497 4.497 0 0 0 8.5 20H9v2h-.5A6.497 6.497 0 0 1 7.2 9.136a8.994 8.994 0 0 1 17.6 0A6.497 6.497 0 0 1 23.5 22Z"/>`),
		g.Group(children),
	)
}