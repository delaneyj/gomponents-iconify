package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Users(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feUsers0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUsers1" fill="currentColor"><path id="feUsers2" d="M8 13a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm8 0a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm-8 2a7.98 7.98 0 0 1 6 2.708V19H2v-1.292A7.98 7.98 0 0 1 8 15Zm8 4v-2.048l-.5-.567a10.057 10.057 0 0 0-1.25-1.193A8.028 8.028 0 0 1 16 15a7.98 7.98 0 0 1 6 2.708V19h-6Z"/></g></g>`),
		g.Group(children),
	)
}