package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeKeyresult(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 0a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4Zm7.25 8a.75.75 0 0 1-.75-.75v-.69L7 10.06v2.19a.75.75 0 0 1-1.5 0v-1.691a.742.742 0 0 1-.059-.059H3.75a.75.75 0 0 1 0-1.5h2.19l3.5-3.5h-.69a.75.75 0 0 1 0-1.5H12v3.25a.75.75 0 0 1-.75.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}