package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tips(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 2v17h-7.586L12 22.414L8.586 19H1V2h22Zm-2 2H3v13h6.414L12 19.586L14.586 17H21V4Z"/>`),
		g.Group(children),
	)
}