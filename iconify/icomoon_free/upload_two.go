package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 14h16v1H0zm16-2v1H0v-1l2-4h4v2h4V8h4zM3.5 5L8 .5L12.5 5H9v4H7V5z"/>`),
		g.Group(children),
	)
}