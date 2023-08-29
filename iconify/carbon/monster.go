package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 22h-2v-2h-2v2h-2v-2h-2v2h-2v-2h-2v2H9v2h2v2h2v-2h2v2h2v-2h2v2h2v-2h2v-2z"/><path fill="currentColor" d="M28 11h-1V4a2.002 2.002 0 0 0-2-2H7a2.002 2.002 0 0 0-2 2v7H4a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h1v4a7.008 7.008 0 0 0 7 7h8a7.008 7.008 0 0 0 7-7v-4h1a2.002 2.002 0 0 0 2-2v-4a2.002 2.002 0 0 0-2-2Zm-3-7v3.382l-2.553-1.276a1 1 0 0 0-.894 0L18 7.881l-3.553-1.776a1 1 0 0 0-.894 0L10 7.881l-3-1.5V4Zm3 13h-3v6a5.006 5.006 0 0 1-5 5h-8a5.006 5.006 0 0 1-5-5v-6H4v-4h3V8.618l2.553 1.277a1 1 0 0 0 .894 0L14 8.118l3.553 1.776a1.001 1.001 0 0 0 .894 0L22 8.119l3 1.5V13h3Z"/><path fill="currentColor" d="M9 14h5v2H9zm9 0h5v2h-5z"/>`),
		g.Group(children),
	)
}