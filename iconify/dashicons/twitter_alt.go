package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitterAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.9 13.8H8.5c-.2 0-.3-.1-.3-.4v-3h5.7c1.1 0 2.1-.9 2.1-2.1c0-1.2-1-2.1-2.1-2.1H8.2V4.1C8.2 2.9 7.2 2 6 2c-1.1 0-2 .9-2 2.1v9.2C4 16 5.5 17.9 8.6 18H14c1.1 0 2.1-1 2.1-2.1c-.1-1.2-1.1-2.1-2.2-2.1z"/>`),
		g.Group(children),
	)
}