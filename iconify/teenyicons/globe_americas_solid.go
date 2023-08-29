package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeAmericasSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM1.197 5.904A6.503 6.503 0 0 0 6 13.826v-.619l-1-1v-1.5l-1-1V8.5a.5.5 0 0 1 .5-.5h4A1.5 1.5 0 0 1 10 9.5v.512c.51.073.915.477.988.988h1.99A6.502 6.502 0 0 0 10 1.498V2.5A1.5 1.5 0 0 1 8.5 4h-1a.5.5 0 0 0-.5.5A1.5 1.5 0 0 1 5.5 6H5v.707l-.44.44a1.5 1.5 0 0 1-2.12 0L1.196 5.903Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}