package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M2 5.077S3.667 2 12 2s10 3.077 10 3.077v13.846S20.333 22 12 22S2 18.923 2 18.923V5.077ZM2 13s3.333 3 10 3s10-3 10-3M2 7s3.333 3 10 3s10-3 10-3"/>`),
		g.Group(children),
	)
}