package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeTask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 12.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm2.28-6.28a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06 0l-1-1a.75.75 0 0 1 1.06-1.06l.47.47l1.97-1.97a.75.75 0 0 1 1.06 0Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M4 0a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4Zm4 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}