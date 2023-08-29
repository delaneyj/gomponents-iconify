package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.926 15.991L13.853 12v7.995L6.926 24zM0 12v7.995l6.926-4.005zM6.926 0v7.995L0 12V4.005zm0 7.995l6.926-3.991V12z"/>`),
		g.Group(children),
	)
}