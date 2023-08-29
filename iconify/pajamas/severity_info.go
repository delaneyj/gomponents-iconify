package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeverityInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 6A6 6 0 1 1 0 6a6 6 0 0 1 12 0ZM6.75 3.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM6 5.25a.75.75 0 0 0-.75.75v2.25a.75.75 0 1 0 1.5 0V6A.75.75 0 0 0 6 5.25Z"/>`),
		g.Group(children),
	)
}