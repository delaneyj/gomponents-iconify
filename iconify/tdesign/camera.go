package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.882 2h8.236l1.5 3H23v16H1V5h5.382l1.5-3Zm1.236 2l-1.5 3H3v12h18V7h-4.618l-1.5-3H9.118ZM12 9.5a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/>`),
		g.Group(children),
	)
}