package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 19.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm2-3h-4v-.129A62 62 0 0 0 8.033.88L8 .75V.5h8v.25l-.033.129A61.999 61.999 0 0 0 14 16.37v.129Z"/>`),
		g.Group(children),
	)
}