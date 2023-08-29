package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Semantics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m2 17l10-5l10 5v4l-10-5l-10 5v-4Zm0-9l10-5l10 5v4L12 7L2 12V8Z"/>`),
		g.Group(children),
	)
}