package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2s3 2 7 2c0 11-7 14-7 14S3 15 3 4c4 0 7-2 7-2z"/>`),
		g.Group(children),
	)
}