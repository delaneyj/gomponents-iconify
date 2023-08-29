package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.472 20.937a1.438 1.438 0 0 1-1.3-.812L10.3 14.343a1.418 1.418 0 0 0-.642-.641l-5.784-2.871a1.462 1.462 0 0 1 .186-2.695l14.952-5a1.46 1.46 0 0 1 1.849 1.847l-5 14.952a1.439 1.439 0 0 1-1.284.994a.525.525 0 0 1-.105.008Zm5.007-16.874a.488.488 0 0 0-.149.024l-14.952 5a.46.46 0 0 0-.058.849l5.78 2.869a2.444 2.444 0 0 1 1.1 1.095l2.87 5.782a.443.443 0 0 0 .445.255a.45.45 0 0 0 .4-.312l5-14.953a.462.462 0 0 0-.433-.607Z"/>`),
		g.Group(children),
	)
}