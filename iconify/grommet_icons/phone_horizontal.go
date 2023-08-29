package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M23 7a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V7Zm-2 5.996V17H3V7h18v4.004a1.038 1.038 0 0 0-.094-.004h-1a1 1 0 0 0 0 2h1c.032 0 .063-.002.094-.004Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}