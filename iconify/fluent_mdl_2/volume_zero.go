package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M677 256h91v1536h-90l-385-384H0V640h293l384-384zm-37 1317V475L347 768H128v512h219l293 293z"/>`),
		g.Group(children),
	)
}