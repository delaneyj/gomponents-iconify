package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4A1.5 1.5 0 0 1 3 2.5h10A1.5 1.5 0 0 1 14.5 4v6a1.5 1.5 0 0 1-1.5 1.5H6.761l-.194.138L4.5 13.1v-1.6H3A1.5 1.5 0 0 1 1.5 10V4ZM3 1a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3v3l1.183-.838L7.24 13H13a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H3Zm8 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM9 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM5 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}