package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 13a2 2 0 0 1 2 2c0 1.11-.89 2-2 2s-2-.89-2-2s.9-2 2-2m8-12c-2.76 0-5 2.24-5 5v2H4c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V10c0-1.1-.9-2-2-2h-1V6c0-1.66 1.34-3 3-3s3 1.34 3 3v2h2V6c0-2.76-2.24-5-5-5m-2 9v10H4V10h12Z"/>`),
		g.Group(children),
	)
}