package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22S3 18 3 5l9-3l9 3c0 13-9 17-9 17ZM4 11h16m-8-9v20"/>`),
		g.Group(children),
	)
}