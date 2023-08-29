package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentTwoDoubleTapSwipeUpTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 16.017a.75.75 0 0 0 .743-.648l.007-.102V4.562l2.22 2.218a.75.75 0 0 0 .976.073l.084-.073a.75.75 0 0 0 .073-.976l-.073-.084l-3.5-3.5a.75.75 0 0 0-.976-.073l-.084.073l-3.5 3.5a.75.75 0 0 0 .976 1.133l.084-.073l2.22-2.218v10.705c0 .414.336.75.75.75Zm0 6a7 7 0 0 0 1.75-13.78v1.564a5.5 5.5 0 1 1-3.5 0V8.237A7 7 0 0 0 12 22.017Zm0-2.5a4.5 4.5 0 0 0 1.751-8.647v1.71a3 3 0 1 1-3.5 0l-.001-1.71A4.501 4.501 0 0 0 12 19.517Z"/>`),
		g.Group(children),
	)
}