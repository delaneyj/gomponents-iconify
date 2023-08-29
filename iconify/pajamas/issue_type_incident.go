package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeIncident(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 4a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4v8a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4zm6.996.165a1.017 1.017 0 1 1 2.012 0L8 11L6.996 4.165zM8 11a1 1 0 1 1 0 2a1 1 0 0 1 0-2z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}