package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListSquareFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 15.75a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-2 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm4-18c0-.69.56-1.25 1.25-1.25h13.5a1.25 1.25 0 1 1 0 2.5h-13.5c-.69 0-1.25-.56-1.25-1.25Zm1.25 6.75a1.25 1.25 0 1 0 0 2.5h13.5a1.25 1.25 0 1 0 0-2.5h-13.5ZM19 31.75c0-.69.56-1.25 1.25-1.25h13.5a1.25 1.25 0 1 1 0 2.5h-13.5c-.69 0-1.25-.56-1.25-1.25ZM12.25 6A6.25 6.25 0 0 0 6 12.25v23.5A6.25 6.25 0 0 0 12.25 42h23.5A6.25 6.25 0 0 0 42 35.75v-23.5A6.25 6.25 0 0 0 35.75 6h-23.5ZM8.5 12.25a3.75 3.75 0 0 1 3.75-3.75h23.5a3.75 3.75 0 0 1 3.75 3.75v23.5a3.75 3.75 0 0 1-3.75 3.75h-23.5a3.75 3.75 0 0 1-3.75-3.75v-23.5Z"/>`),
		g.Group(children),
	)
}