package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 17c0-5 1-11 8-11s8 6 8 11s3 8 3 8H5s3-3 3-8Zm12 8s0 4-4 4s-4-4-4-4m4-22v3"/>`),
		g.Group(children),
	)
}