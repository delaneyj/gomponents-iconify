package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardTextThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 152a4 4 0 0 1-4 4H96a4 4 0 0 1 0-8h64a4 4 0 0 1 4 4Zm-4-36H96a4 4 0 0 0 0 8h64a4 4 0 0 0 0-8Zm52-68v168a12 12 0 0 1-12 12H56a12 12 0 0 1-12-12V48a12 12 0 0 1 12-12h38.08a44 44 0 0 1 67.84 0H200a12 12 0 0 1 12 12ZM92 64v4h72v-4a36 36 0 0 0-72 0Zm112-16a4 4 0 0 0-4-4h-32.83A43.71 43.71 0 0 1 172 64v8a4 4 0 0 1-4 4H88a4 4 0 0 1-4-4v-8a43.71 43.71 0 0 1 4.83-20H56a4 4 0 0 0-4 4v168a4 4 0 0 0 4 4h144a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}