package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PavedRoadAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 6a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H15Zm8 5a1 1 0 1 1 2 0v4a1 1 0 0 1-2 0v-4Zm0 11a1 1 0 1 1 2 0v4a1 1 0 0 1-2 0v-4Zm1 10a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0v-4a1 1 0 0 0-1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}