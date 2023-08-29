package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.586 3.586a2 2 0 1 1 2.828 2.828l-.793.793l-2.828-2.828l.793-.793Zm-2.207 2.207L3 14.172V17h2.828l8.38-8.379l-2.83-2.828Z"/>`),
		g.Group(children),
	)
}