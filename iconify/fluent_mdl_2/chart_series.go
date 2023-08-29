package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartSeries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M512 640h384v1280H0V384h512v256zM384 512H128v1280h256V512zm128 256v1024h256V768H512zm1536 128v1024h-896V128h512v768h384zm-512-640h-256v1536h256V256zm384 768h-256v768h256v-768z"/>`),
		g.Group(children),
	)
}