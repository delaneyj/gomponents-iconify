package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taurus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#e26a3b"/><path fill="#fff" stroke="#fff" stroke-miterlimit="10" d="M46.6 17c-2.3 0-5 .4-6.8 1.9c-2.6 2.1-4.1 6.8-7.8 7.1c-3.7-.3-5.2-4.9-7.8-7.1c-1.8-1.5-4.5-1.8-6.8-1.9c-1.9 0-1.9 3 0 3c3.8 0 5.9 1.7 8 4.8c.6.9 1.3 1.7 1.9 2.3c-3.3 1.8-5.4 5.6-5.6 9.4c-.3 5.9 4.9 10.2 10.2 10.5c5.6.3 9.7-4.7 10.2-10c0-.1.1-.3 0-.4v-.4c-.3-4-2.5-7.2-5.7-8.8c.8-.7 1.4-1.5 2.1-2.5c2.1-3.1 4.2-4.7 8-4.8c2-.1 2-3.1.1-3.1zm-7.3 19.5c-.5 9.7-14.1 9.7-14.6 0c-.6-9.7 14.1-9.6 14.6 0z"/>`),
		g.Group(children),
	)
}