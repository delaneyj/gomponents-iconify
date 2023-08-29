package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiStickerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.902 10.598a9.986 9.986 0 0 0-9.381 3.873a4.977 4.977 0 0 1-3.854-1.246l-1.334 1.49a6.976 6.976 0 0 0 4.014 1.753A9.969 9.969 0 0 0 10.5 20.5c0 .476.033.944.098 1.402C5.738 21.221 2 17.047 2 12C2 6.477 6.477 2 12 2c5.047 0 9.22 3.739 9.902 8.598Zm-.031 2.019a7.99 7.99 0 0 0-7.964 3.35A7.958 7.958 0 0 0 12.5 20.5c0 .467.04.925.117 1.37l9.254-9.253ZM8.5 11.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm7 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}