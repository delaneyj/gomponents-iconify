package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 7H5a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2zm-9.684 7.316l1.602 1.4c.305.266.691.398 1.082.398s.777-.133 1.082-.398l1.602-1.4l-.037.037l3.646 3.646H5.707l3.646-3.646l-.037-.037zM5 17.293V10.54l3.602 3.151L5 17.293zm10.398-3.602L19 10.54v6.75l-3.602-3.599zM19 9v.21l-6.576 5.754a.68.68 0 0 1-.848 0L5 9.21V9h14z"/>`),
		g.Group(children),
	)
}