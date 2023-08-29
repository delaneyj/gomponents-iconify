package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrescentMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M43.139 2a29.885 29.885 0 0 1 5.121 16.756c0 16.701-13.686 30.24-30.57 30.24a30.656 30.656 0 0 1-15.689-4.285C7.209 54.963 17.93 62 30.318 62C47.816 62 62 47.969 62 30.66C62 17.867 54.246 6.871 43.139 2z"/>`),
		g.Group(children),
	)
}