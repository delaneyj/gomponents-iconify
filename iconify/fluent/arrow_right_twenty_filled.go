package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.265 3.205a.75.75 0 0 0-1.03 1.09l5.239 4.955H2.75a.75.75 0 0 0 0 1.5h12.726l-5.241 4.957a.75.75 0 1 0 1.03 1.09l6.418-6.07a.995.995 0 0 0 .3-.566a.753.753 0 0 0-.002-.329a.995.995 0 0 0-.298-.557l-6.418-6.07Z"/>`),
		g.Group(children),
	)
}