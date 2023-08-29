package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayReverseResume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 256v1536h-128V256h128zM192 1024l1088-768v1536L192 1024zm960-521l-738 521l738 521V503z"/>`),
		g.Group(children),
	)
}