package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#d8544a"/><path fill="#fff" stroke="#fff" stroke-miterlimit="10" d="M44.3 17c-4.3-.2-6.8 3.2-7.9 7.1c-1.5 5.3-2.9 10.6-4.4 15.9c-1.5-5.3-2.9-10.5-4.4-15.8c-1.1-3.9-3.6-7.4-7.9-7.1c-3.9.2-6.5 3.8-6.7 7.5c-.1 2 2.8 1.9 2.9 0c.1-2.6 1.9-4.4 4.4-4.5c2.6-.1 3.9 2.9 4.4 4.9c1.9 7 3.9 14 5.8 20.9c.4 1.5 2.4 1.5 2.8 0c1.9-7 3.9-14 5.8-20.9c.7-2.3 1.8-4.8 4.5-4.9c2.5-.2 4.3 2.2 4.4 4.5c.1 1.9 3 2 2.9 0c-.1-3.8-2.7-7.4-6.6-7.6z"/>`),
		g.Group(children),
	)
}