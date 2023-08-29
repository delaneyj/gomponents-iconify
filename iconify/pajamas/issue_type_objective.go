package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeObjective(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 0a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4Zm6.75 2a.75.75 0 0 1 .75.75v1.691a.756.756 0 0 1 .059.059h1.691a.75.75 0 0 1 0 1.5h-2.19L9.45 7.611A1.502 1.502 0 0 1 8 9.5a1.5 1.5 0 1 1 .389-2.95L10 4.94V2.75a.75.75 0 0 1 .75-.75ZM8.584 3.724a.75.75 0 0 1-.72.779a3.5 3.5 0 1 0 3.633 3.631a.75.75 0 0 1 1.5.058a5 5 0 1 1-5.191-5.188a.75.75 0 0 1 .778.72Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}