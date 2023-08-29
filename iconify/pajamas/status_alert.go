package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 6A6 6 0 1 1 0 6a6 6 0 0 1 12 0zM5 3a1 1 0 0 1 2 0v3a1 1 0 0 1-2 0V3zm1 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}