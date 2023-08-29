package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M23 21.8V16h20v4.5c-1.7 1.7-3.3 4.2-5 7.4c-1.7 3.2-3 6.7-3.9 10.3c-.9 3.6-1.3 6.9-1.3 9.7h-5.6c.1-4.5 1-9.1 2.6-13.7c1.6-4.7 3.8-8.8 6.6-12.5l-13.4.1"/>`),
		g.Group(children),
	)
}