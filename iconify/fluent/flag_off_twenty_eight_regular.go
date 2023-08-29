package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOffTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m2.397 1.554l.073-.084a.75.75 0 0 1 .976-.073l.084.073l20 20a.75.75 0 0 1-.976 1.133l-.084-.073L17.938 18H5.5v6.25a.75.75 0 0 1-1.493.102L4 24.25V4.061l-1.53-1.53a.75.75 0 0 1-.073-.977l.073-.084l-.073.084ZM7.06 2.999L23.25 3a.75.75 0 0 1 .683 1.06l-.048.09l-3.999 6.35l3.999 6.35a.75.75 0 0 1-.533 1.143L23.25 18h-1.189l-1.5-1.5h1.33l-3.526-5.6a.75.75 0 0 1-.05-.705l.05-.095l3.527-5.6H8.56L7.06 3ZM5.5 5.561V16.5l10.938-.001L5.5 5.561Z"/>`),
		g.Group(children),
	)
}