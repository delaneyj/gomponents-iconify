package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsStopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 5c-4.411 0-8 3.589-8 8s3.589 8 8 8s8-3.589 8-8s-3.589-8-8-8zm1 8h-2V8h2v5zM9 2h6v2H9zm9.707 2.293l2 2l-1.414 1.414l-2-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}