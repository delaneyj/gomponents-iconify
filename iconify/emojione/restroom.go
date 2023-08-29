package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Restroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M46 23h-8c-1 0-2 1-2 2l-3 11h3l2-9v4l-2 9h2v11h3V40h2v11h3V40h2l-2-9v-4l2 9h3l-3-11c0-1-1-2-2-2m0-4c0 1-1 2-2 2h-4c-1 0-2-1-2-2v-4c0-1 1-2 2-2h4c1 0 2 1 2 2v4m-18 4H16c-1 0-2 1-2 2l-1 11h3l1-9l1 24h3l1-13l1 13h3l1-24l1 9h3l-1-11c0-1-1-2-2-2m-2-4c0 1-1 2-2 2h-4c-1 0-2-1-2-2v-4c0-1 1-2 2-2h4c1 0 2 1 2 2v4"/>`),
		g.Group(children),
	)
}