package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellSlashO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m15.2 0l-3.6 3.6C11.2 3 10.4 2.3 9 2.1V1s.1-1-1-1s-1 1-1 1v1.1C4 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v1.3l-2 2v.7h.7L16 .6V0h-.8zM5 10.3c0-.1 0-.1 0 0V5c0-.1.1-1.6 2.2-1.9l.8-.2l.8.1c1.2.2 1.8.8 2 1.3l-5.8 6zm7-.1V5.6l-1 1v3.5c0 .9.5 1.7 1.3 2.1l.7.4v.4H4.7l-1 1h2.4s-.1 2 2 2s2-2 2-2H14v-2l-1.3-.6c-.4-.3-.7-.7-.7-1.2z"/>`),
		g.Group(children),
	)
}