package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PmTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M485.317 146.623H26.682c-35.576 0-35.576-59.374 0-59.374h458.636c35.576 0 35.576 59.374 0 59.374zm0 139.064H26.682c-35.576 0-35.576-59.375 0-59.375h458.636c35.576 0 35.576 59.375 0 59.375zm0 139.064H26.682c-35.576 0-35.576-59.374 0-59.374h458.636c35.576 0 35.576 59.374 0 59.374z"/>`),
		g.Group(children),
	)
}