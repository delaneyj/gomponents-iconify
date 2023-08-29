package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m253 281l-16 44l-66-66L64 365l-30-30l108-107q-40-44-63-97h42q20 39 50 71q45-50 67-114H0V45h149V3h43v42h149v43h-62q-24 78-79 139l-1 1zm120-108l96 256h-42l-24-64H301l-24 64h-42l96-256h42zm-56 150h70l-35-93z"/>`),
		g.Group(children),
	)
}