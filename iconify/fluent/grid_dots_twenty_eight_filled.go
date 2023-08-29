package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridDotsTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 20a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm8.5 0a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm-17 0a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm8.5-8.5a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm8.5 0a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm-17 0a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5ZM14 3a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm8.5 0a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm-17 0a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Z"/>`),
		g.Group(children),
	)
}