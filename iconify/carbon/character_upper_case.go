package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CharacterUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 23h-6a2.002 2.002 0 0 1-2-2V11a2.002 2.002 0 0 1 2-2h6v2h-6v10h6zM20 12a3.003 3.003 0 0 0-3-3h-5v14h5a3.003 3.003 0 0 0 3-3v-2a2.977 2.977 0 0 0-.78-2a2.977 2.977 0 0 0 .78-2zm-6-1h3a1.001 1.001 0 0 1 1 1v2a1.001 1.001 0 0 1-1 1h-3zm4 9a1 1 0 0 1-1 1h-3v-4h3a1 1 0 0 1 1 1zM8 9H4a2.002 2.002 0 0 0-2 2v12h2v-5h4v5h2V11a2.002 2.002 0 0 0-2-2zm-4 7v-5h4v5z"/>`),
		g.Group(children),
	)
}