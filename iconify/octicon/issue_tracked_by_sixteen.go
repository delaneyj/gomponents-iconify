package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTrackedBySixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1.5 8a6.5 6.5 0 0 1 13 0A.75.75 0 0 0 16 8a8 8 0 1 0-8 8a.75.75 0 0 0 0-1.5A6.5 6.5 0 0 1 1.5 8Z"/><path fill="currentColor" d="M8 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm3.573 5.823l-2.896-2.896a.25.25 0 0 1 0-.354l2.896-2.896a.25.25 0 0 1 .427.177V11.5h3.25a.75.75 0 0 1 0 1.5H12v2.146a.25.25 0 0 1-.427.177Z"/>`),
		g.Group(children),
	)
}