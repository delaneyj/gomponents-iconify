package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M60.571 24.301a2.604 2.604 0 0 0-2.604-2.604h-4.594a2.598 2.598 0 0 0-2.59 2.463l-.014-.001l-11.276 50.978l-.015.066l-.011.048h.006a2.55 2.55 0 0 0-.045.449a2.595 2.595 0 0 0 2.406 2.584v.02h4.792a2.595 2.595 0 0 0 2.577-2.336l.013.001l11.257-50.972l-.008-.001a2.58 2.58 0 0 0 .106-.695z"/>`),
		g.Group(children),
	)
}