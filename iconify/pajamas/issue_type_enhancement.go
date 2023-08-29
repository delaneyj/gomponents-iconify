package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeEnhancement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 0a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4zm8 9.25V4H6.75a.75.75 0 0 0 0 1.5h2.69l-5.22 5.22a.75.75 0 1 0 1.06 1.06l5.22-5.22v2.69a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}