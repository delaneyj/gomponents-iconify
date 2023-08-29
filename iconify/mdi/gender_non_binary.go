package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderNonBinary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3h-2v2.27L9.04 4.13l-1 1.74L10 7L8.04 8.13l1 1.74L11 8.73v3.37a5 5 0 1 0 2 0V8.73l1.96 1.14l1-1.74L14 7l1.96-1.13l-1-1.74L13 5.27V3m-1 17c-1.65 0-3-1.35-3-3s1.35-3 3-3s3 1.35 3 3s-1.35 3-3 3Z"/>`),
		g.Group(children),
	)
}