package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeFeatureFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 0a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4zm.225 4.803L4 5v6.5a.5.5 0 0 0 1 0V9.303A3.177 3.177 0 0 1 8 9.5a3.176 3.176 0 0 0 3.775-.303L12 9V4l-.225.197A3.176 3.176 0 0 1 8 4.5a3.176 3.176 0 0 0-3.775.303z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}