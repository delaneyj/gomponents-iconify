package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gravefour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 1024H0q0-47 56.5-87.5T210 869l-82-549q0-30 82.5-109T387 66T512 0t125 66t176.5 145T896 320l-82 549q97 27 153.5 67.5t56.5 87.5z"/>`),
		g.Group(children),
	)
}