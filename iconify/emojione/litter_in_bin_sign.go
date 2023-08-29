package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LitterInBinSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><g fill="#fff"><path d="m28 38l-1 10h-8l-1-10h-2v13h14V38z"/><path d="m20 37l3-3l3 4l-4 2zm25-14H33c-1 0-2 1-2 2l-1 4l-5 3l1 2l6-2l2-5l1 24h3l1-13l1 13h3l1-24l1 9h3l-1-11c0-1-1-2-2-2m-2-4c0 1-1 2-2 2h-4c-1 0-2-1-2-2v-4c0-1 1-2 2-2h4c1 0 2 1 2 2v4"/></g>`),
		g.Group(children),
	)
}