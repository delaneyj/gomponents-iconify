package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Podio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.2 13.8L9.1 9.7c-.2-.2-.4-.5-.4-.9c0-.7.6-1.3 1.3-1.3c.1 0 .3 0 .4.1c.7.2 1.1.9.9 1.6c-.1.5.1 1 .6 1.1c.5.1 1-.1 1.1-.6c.1-.3.1-.6.1-.9c0-1.7-1.4-3.1-3.1-3.1c-.8 0-1.6.3-2.2.9c-1.2 1.2-1.2 3.2 0 4.4l4.1 4.1c.4.4.9.4 1.3 0c.4-.3.4-.9 0-1.3zM14.9 4C12.2 1.4 8 1.3 5.3 4c-2.7 2.7-2.7 7 0 9.6l4.1 4.1c.4.4.9.4 1.3 0s.4-.9 0-1.3l-4.1-4.1c-1.9-1.9-1.9-5.1 0-7c1.9-1.9 5.1-1.9 7 0c1.6 1.7 1.8 4.4.5 6.3c-.3.4-.1 1 .3 1.3c.4.3 1 .1 1.3-.3c1.7-2.7 1.4-6.3-.8-8.6z"/>`),
		g.Group(children),
	)
}