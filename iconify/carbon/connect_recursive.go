package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectRecursive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 7H13.828l3.586-3.586L16 2l-6 6l6 6l1.414-1.414L13.828 9H28v11H11.899A5.014 5.014 0 0 0 8 16.101V2H6v14.101A5 5 0 0 0 6 25.9V30h2v-4.101A5.014 5.014 0 0 0 11.899 22H28a2.002 2.002 0 0 0 2-2V9a2.002 2.002 0 0 0-2-2ZM7 24a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3Z"/>`),
		g.Group(children),
	)
}