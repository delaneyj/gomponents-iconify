package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickFromTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M336 113.375H176v144H56V296l199.8 200L456 296.007v-38.632H336Zm81.361 176L255.826 450.746L94.616 289.375H208v-144h96v144ZM56 17.376h400v32H56z"/>`),
		g.Group(children),
	)
}