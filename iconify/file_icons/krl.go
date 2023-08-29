package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Krl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0v512h512V0H0zm11.152 494.296l170.55-252.647L20.672 16.964h96.02L276.68 240.205l-171.524 254.09H11.152zm394.346 0l-139.505-206.66l33.316-47.73l-33.494-43.954L393.962 16.964h96.018L328.952 241.649l170.549 252.647h-94.003z"/>`),
		g.Group(children),
	)
}