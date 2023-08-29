package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignLeftRotateNinetySixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.25 1a.75.75 0 0 1 .75.75v8.5a.75.75 0 0 1-1.5 0v-8.5a.75.75 0 0 1 .75-.75Zm-4 0a.75.75 0 0 1 .75.75v12.5a.75.75 0 0 1-1.5 0V1.75A.75.75 0 0 1 8.25 1Zm-4 0a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0v-4.5A.75.75 0 0 1 4.25 1Z"/>`),
		g.Group(children),
	)
}