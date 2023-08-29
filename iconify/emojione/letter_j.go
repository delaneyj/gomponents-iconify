package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterJ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M28.6 35.2v.7c0 2.2.3 3.8.7 4.7c.4.9 1.3 1.4 2.7 1.4c1.3 0 2.2-.5 2.7-1.5c.3-.6.4-1.6.4-3V17.1h6.1v20.2c0 2.5-.4 4.4-1.3 5.9c-1.4 2.4-4 3.7-7.6 3.7s-6.2-1-7.4-3c-1.3-2-1.9-4.7-1.9-8.1v-.7h5.6z"/>`),
		g.Group(children),
	)
}