package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.002 2.999a1 1 0 0 1 1.6-.8l7.998 6c.768.578.36 1.8-.6 1.8H9.053a1 1 0 0 0-.793.392l-2.466 3.214c-.581.758-1.793.347-1.793-.609V3Zm8.997 6L5.002 3v9.997l2.465-3.213A2 2 0 0 1 9.054 9H13Z"/>`),
		g.Group(children),
	)
}