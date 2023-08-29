package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiStickerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.52 19.863a9.948 9.948 0 0 1 .826-3.395a6.977 6.977 0 0 1-4.013-1.753l1.334-1.49a4.977 4.977 0 0 0 3.854 1.246a9.987 9.987 0 0 1 7.342-3.951a8 8 0 1 0-9.343 9.343Zm8.503-7.227a8.008 8.008 0 0 0-6.387 6.387l6.387-6.387ZM22 12c0 .168-.004.334-.012.5L12.5 21.988A10.11 10.11 0 0 1 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Zm-12-2a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm7 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}