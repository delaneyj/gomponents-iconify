package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 18h-7v-4h3a2.002 2.002 0 0 0 2-2V8a2.002 2.002 0 0 0-2-2h-3V2h-2v4h-3a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h3v4H8a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h7v4h2v-4h7a2.002 2.002 0 0 0 2-2v-4a2.002 2.002 0 0 0-2-2ZM12 8h8v4h-8Zm12 16H8v-4h16Z"/>`),
		g.Group(children),
	)
}