package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElipsisV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feElipsisV0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feElipsisV1" fill="currentColor"><path id="feElipsisV2" d="M12 20a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/></g></g>`),
		g.Group(children),
	)
}