package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitRepositoryLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21v2.5l-3-2l-3 2V21h-.5A3.5 3.5 0 0 1 3 17.5V5a3 3 0 0 1 3-3h14a1 1 0 0 1 1 1v17a1 1 0 0 1-1 1h-7Zm0-2h6v-3H6.5a1.5 1.5 0 0 0 0 3H7v-2h6v2Zm6-5V4H6v10.035c.163-.023.33-.035.5-.035H19ZM7 5h2v2H7V5Zm0 3h2v2H7V8Zm0 3h2v2H7v-2Z"/>`),
		g.Group(children),
	)
}