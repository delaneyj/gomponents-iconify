package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Film(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H4Zm3 2h6v4H7V5Zm8 8v2h1v-2h-1Zm-2-2H7v4h6v-4Zm2 0h1V9h-1v2Zm1-4V5h-1v2h1ZM5 5v2H4V5h1Zm0 4H4v2h1V9Zm-1 4h1v2H4v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}