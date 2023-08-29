package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropPhotoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="19" cy="22" r="2" fill="currentColor"/><path fill="currentColor" d="M27 16h-3.28l-.543-1.633A1.998 1.998 0 0 0 21.279 13h-4.558a1.998 1.998 0 0 0-1.898 1.368L14.28 16H11a2.002 2.002 0 0 0-2 2v9a2.002 2.002 0 0 0 2 2h16a2.002 2.002 0 0 0 2-2v-9a2.002 2.002 0 0 0-2-2Zm-8 10a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/><path fill="none" d="M19 26a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2Z"/><path fill="currentColor" d="M27 11h2v2h-2zm0-4h2v2h-2zm0-4h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zM7 3h2v2H7zM3 3h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3z"/>`),
		g.Group(children),
	)
}