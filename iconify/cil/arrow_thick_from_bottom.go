package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickFromBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256.2 16L56 215.993v38.632h120v144h160v-144h120V216ZM304 222.625v144h-96v-144H94.639L256.174 61.254l161.21 161.371Zm-248 240h400v32H56z"/>`),
		g.Group(children),
	)
}