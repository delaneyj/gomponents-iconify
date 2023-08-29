package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M23.3 20.4c2.3-2.4 5.3-3.6 8.9-3.6c4.8 0 8.4 1.6 10.6 4.8c1.2 1.8 1.9 3.6 2 5.4h-6.1c-.4-1.4-.9-2.4-1.5-3.1c-1.1-1.3-2.7-1.9-4.8-1.9c-2.2 0-3.9.9-5.1 2.7c-1.2 1.8-1.9 4.3-1.9 7.5c0 3.3.7 5.7 2 7.3c1.3 1.6 3 2.4 5 2.4c2.1 0 3.7-.7 4.8-2.1c.6-.7 1.1-1.9 1.5-3.4h6c-.5 3.2-1.8 5.7-4 7.7c-2.1 2-4.9 3-8.2 3c-4.1 0-7.4-1.3-9.7-4c-2.4-2.7-3.5-6.4-3.5-11.1c-.1-5 1.3-8.9 4-11.6"/>`),
		g.Group(children),
	)
}