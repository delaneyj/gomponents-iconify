package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taskboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v1792H0V128h1920zm-128 128h-768v1024h768V256zm-1664 0v384h768V256H128zm0 1536h768V768H128v1024zm1664 0v-384h-768v384h768z"/>`),
		g.Group(children),
	)
}