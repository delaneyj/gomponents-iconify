package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func X(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M519 128L315 380l222 274h-93L269 437L93 654H0l223-274L19 128h93l157 194l156-194h94z"/>`),
		g.Group(children),
	)
}