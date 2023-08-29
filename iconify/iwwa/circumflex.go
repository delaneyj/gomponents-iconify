package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circumflex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M6.419 27.287L20 11.162l.113.113l13.475 16.019a.847.847 0 0 0 1.18-.013a.825.825 0 0 0 .233-.587a.838.838 0 0 0-.238-.595L20.647 9.444a.81.81 0 0 0-.628-.225a.843.843 0 0 0-.67.23L5.22 26.118c-.295.313-.299.844.025 1.18a.846.846 0 0 0 1.174-.011z"/>`),
		g.Group(children),
	)
}