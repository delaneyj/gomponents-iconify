package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchDeleted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.28 1.22a.75.75 0 1 0-1.06 1.06l1.22 1.22l-1.22 1.22a.75.75 0 0 0 1.06 1.06l1.22-1.22l1.22 1.22a.75.75 0 1 0 1.06-1.06L12.56 3.5l1.22-1.22a.75.75 0 0 0-1.06-1.06L11.5 2.44l-1.22-1.22Zm-2.5 6a.75.75 0 0 1 0 1.06l-2.045 2.046A2.499 2.499 0 0 1 4.5 15a2.5 2.5 0 0 1-.75-4.886V5.886a2.501 2.501 0 1 1 1.5 0v2.803l1.47-1.47a.75.75 0 0 1 1.06 0ZM5.5 12.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0-9a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}