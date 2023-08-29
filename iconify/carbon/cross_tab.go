package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossTab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H10a2.002 2.002 0 0 0-2 2v3H4a2.002 2.002 0 0 0-2 2v19a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2ZM10 4h18v3H10Zm18 10h-8V9h8Zm-18 7v-5h8v5Zm8 2v5h-8v-5ZM8 21H4v-5h4ZM18 9v5h-8V9Zm2 7h8v5h-8ZM8 9v5H4V9ZM4 23h4v5H4Zm16 5v-5h8v5Z"/>`),
		g.Group(children),
	)
}