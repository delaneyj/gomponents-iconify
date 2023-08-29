package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellWithSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ff5a79"/><circle cx="32" cy="32" r="24" fill="#333"/><path fill="#fff" d="M42 32c-1.8-3.1-1.5-5.9-1.5-8c0-4.4-2.5-7.1-6.1-8.2c0-1.4-.8-2.6-2.4-2.6c-1.6 0-2.4 1.1-2.4 2.6c-3.6 1.1-6.1 3.8-6.1 8.2c0 2.1.3 4.9-1.5 8c-1.2 2.2-3.9 2.7-3.9 9.1H46c-.1-6.5-2.7-6.9-4-9.1m-9.2 10.9c0 2.4-2 4.4-4.4 4.4c-2.4 0-4.4-2-4.4-4.4h8.8"/><path fill="#ff5a79" d="m9.23 13.474l4.243-4.242l41.294 41.294l-4.242 4.243z"/>`),
		g.Group(children),
	)
}