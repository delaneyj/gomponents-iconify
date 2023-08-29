package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Call(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m308 288l-88 87q18 65 95 160.5T487.5 708T649 803l87-87q15-15 53-12.5t76.5 12T929 737q56 27 82 66.5t3 61.5L907 972q-43 44-117 49.5t-160-23T448 908T265 759T115.5 575.5T24.5 393t-23-159.5T51 117L158 10q21-21 55 4.5T286 95q19 27 30 97t-8 96z"/>`),
		g.Group(children),
	)
}