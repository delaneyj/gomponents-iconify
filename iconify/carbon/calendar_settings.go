package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 23v-2h-2.09a5.96 5.96 0 0 0-1.024-2.47l1.478-1.48l-1.414-1.414l-1.479 1.479A5.958 5.958 0 0 0 23 16.09V14h-2v2.09a5.958 5.958 0 0 0-2.47 1.024l-1.48-1.478l-1.414 1.414l1.479 1.479A5.962 5.962 0 0 0 16.09 21H14v2h2.09a5.962 5.962 0 0 0 1.024 2.47l-1.478 1.48l1.414 1.414l1.479-1.479A5.958 5.958 0 0 0 21 27.91V30h2v-2.09a5.958 5.958 0 0 0 2.47-1.024l1.48 1.478l1.414-1.414l-1.479-1.479A5.96 5.96 0 0 0 27.91 23Zm-8 3a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/><path fill="currentColor" d="M28 6a2 2 0 0 0-2-2h-4V2h-2v2h-8V2h-2v2H6a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h4v-2H6V6h4v2h2V6h8v2h2V6h4v6h2Z"/>`),
		g.Group(children),
	)
}