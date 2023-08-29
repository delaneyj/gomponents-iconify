package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodHappySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 20a10 10 0 1 1 0-20a10 10 0 0 1 0 20zM6.5 9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm7 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm2.16 3H4.34a6 6 0 0 0 11.32 0z"/>`),
		g.Group(children),
	)
}