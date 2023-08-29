package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M37.7 18.1c2 .7 3.7 1.9 4.9 3.7c1 1.4 1.7 3 2.1 4.7s.6 3.3.6 4.8c0 3.9-.8 7.1-2.3 9.8c-2.1 3.6-5.3 5.4-9.7 5.4H20.7v-29h12.5c1.8 0 3.3.2 4.5.6m-11.1 4.4v18.9h5.6c2.9 0 4.9-1.4 6-4.2c.6-1.5.9-3.4.9-5.5c0-3-.5-5.2-1.4-6.8c-.9-1.6-2.8-2.4-5.5-2.4h-5.6"/>`),
		g.Group(children),
	)
}