package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.75 9.25a.75.75 0 0 1 0 1.5h-5.5a.75.75 0 0 1 0-1.5h5.5Zm4.5.75a7.25 7.25 0 1 0-2.681 5.63l4.9 4.9l.085.073a.75.75 0 0 0 .976-1.133l-4.9-4.901A7.22 7.22 0 0 0 17.25 10Zm-13 0a5.75 5.75 0 1 1 11.5 0a5.75 5.75 0 0 1-11.5 0Z"/>`),
		g.Group(children),
	)
}