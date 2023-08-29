package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hashtag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m959 896l64-256H769l-64 256h254zm768-504l-56 224q-7 24-31 24h-327l-64 256h311q15 0 25 12q10 14 6 28l-56 224q-5 24-31 24h-327l-81 328q-7 24-31 24H841q-16 0-26-12q-9-12-6-28l78-312H633l-81 328q-7 24-31 24H296q-15 0-25-12q-9-12-6-28l78-312H32q-15 0-25-12q-9-12-6-28l56-224q7-24 31-24h327l64-256H168q-15 0-25-12q-10-14-6-28l56-224q5-24 31-24h327l81-328q7-24 32-24h224q15 0 25 12q9 12 6 28l-78 312h254l81-328q7-24 32-24h224q15 0 25 12q9 12 6 28l-78 312h311q15 0 25 12q9 12 6 28z"/>`),
		g.Group(children),
	)
}