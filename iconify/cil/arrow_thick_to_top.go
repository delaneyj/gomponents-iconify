package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickToTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M56 297.365V336h120v160h160V336h120v-38.626l-199.8-200ZM304 304v160h-96V304H94.639l161.535-161.37L417.384 304ZM56 16.002h400v32H56z"/>`),
		g.Group(children),
	)
}