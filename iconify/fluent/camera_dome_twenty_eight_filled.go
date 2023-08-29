package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraDomeTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.75C2 3.784 2.784 3 3.75 3h20.5a1.75 1.75 0 1 1 0 3.5H3.75A1.75 1.75 0 0 1 2 4.75ZM9 16.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Zm5 3.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7ZM3 8h22v7c0 6.075-4.925 11-11 11S3 21.075 3 15V8Zm4.5 8.5a6.5 6.5 0 1 0 13 0a6.5 6.5 0 0 0-13 0Z"/>`),
		g.Group(children),
	)
}