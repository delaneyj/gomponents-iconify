package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.7 11.4c-.5-.2-.7-.7-.7-1.2V5s0-2.4-3-2.9V1s.1-1-1-1s-1 1-1 1v1.1C4 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v2h4s-.1 2 2 2s2-2 2-2h4v-2l-1.3-.6zM13 13H3v-.4l.7-.4c.8-.3 1.3-1.1 1.3-2V5c0-.1 0-1.6 2.2-1.9l.8-.2l.8.1c2 .4 2.2 1.7 2.2 2v5.2c0 .9.5 1.7 1.3 2.1l.7.4v.3z"/>`),
		g.Group(children),
	)
}