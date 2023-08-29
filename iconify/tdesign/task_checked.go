package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v11h-2V5h-2v2H7V5H5v16h6v2H3V3h4V1Zm2 4h6V3H9v2Zm14.657 11.586l-7.07 7.07l-4.243-4.242L13.758 18l2.828 2.828l5.657-5.656l1.414 1.414Z"/>`),
		g.Group(children),
	)
}