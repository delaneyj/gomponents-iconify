package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 11.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM5 10a5 5 0 1 1 10 0a5 5 0 0 1-10 0Zm5-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-8 4a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm8-7a7 7 0 1 0 0 14a7 7 0 0 0 0-14Z"/>`),
		g.Group(children),
	)
}