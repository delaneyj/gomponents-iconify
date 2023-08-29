package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickToBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M336 176.005V16H176v160H56v38.623l199.8 200L456 214.637v-38.632Zm-80.174 193.371L94.616 208.005H208V48h96v160h113.361ZM56 464h400v32H56z"/>`),
		g.Group(children),
	)
}