package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrowdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M467 1004L19 553Q0 534 0 507t19-45l90-92q19-19 45.5-19t45.5 19l184 185V64q0-26 18.5-45T448 0h127q27 0 45.5 19T639 64v492l185-186q19-19 45.5-19t45.5 19l90 92q19 18 19 45t-19 46l-448 451q-18 19-45 19t-45-19z"/>`),
		g.Group(children),
	)
}