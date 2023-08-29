package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeSails(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#28a3b2" d="M4.6 30S-6.108 12.967 15.925 2v28H4.6m14.737 0V12.645S22.853 18.381 30 30H19.337"/>`),
		g.Group(children),
	)
}