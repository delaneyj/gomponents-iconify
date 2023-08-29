package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugDroidLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m189 43.47l15.22-15.23a6 6 0 0 0-8.48-8.48L180 35.54a85.76 85.76 0 0 0-104 0L60.24 19.76a6 6 0 0 0-8.48 8.48L67 43.47A85.7 85.7 0 0 0 42 104v40a86 86 0 0 0 172 0v-40a85.7 85.7 0 0 0-25-60.53ZM128 30a74.09 74.09 0 0 1 74 74v10H54v-10a74.09 74.09 0 0 1 74-74Zm0 188a74.09 74.09 0 0 1-74-74v-18h148v18a74.09 74.09 0 0 1-74 74Zm18-134a10 10 0 1 1 10 10a10 10 0 0 1-10-10Zm-56 0a10 10 0 1 1 10 10a10 10 0 0 1-10-10Z"/>`),
		g.Group(children),
	)
}