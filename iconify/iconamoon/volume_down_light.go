package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linejoin="round" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32.2 32.2 0 0 0 13 12Z"/><path stroke-linecap="round" d="M16.877 9a36.485 36.485 0 0 1 0 6"/></g>`),
		g.Group(children),
	)
}