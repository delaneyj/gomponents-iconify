package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrescentMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffce31" d="M43.1 2c3.2 4.8 5.1 10.6 5.1 16.8C48.3 35.5 34.6 49 17.7 49C12 49 6.6 47.4 2 44.7C7.2 55 17.9 62 30.3 62C47.8 62 62 48 62 30.7C62 17.9 54.2 6.9 43.1 2z"/>`),
		g.Group(children),
	)
}