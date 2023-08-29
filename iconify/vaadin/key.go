package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m8.1 7l-.6-.3L15 0h-2L6 6.1C5.7 6 5.4 6 5 6c-2.8 0-5 2.2-5 5s2.3 5 5 5s5-2.2 5-5c0-.6-.1-1.2-.3-1.7L11 8V6h2V4h2l1-1V0L8.1 7zM4 13.2c-.7 0-1.2-.6-1.2-1.2s.6-1.2 1.2-1.2s1.2.6 1.2 1.2s-.5 1.2-1.2 1.2z"/>`),
		g.Group(children),
	)
}