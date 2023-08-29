package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushRules(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.25 4.07V1.75a.75.75 0 0 1 1.5 0v2.32a4.001 4.001 0 0 1 0 7.86v2.32a.75.75 0 0 1-1.5 0v-2.32a4.001 4.001 0 0 1 0-7.86ZM4 10.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM9.75 4a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5h-5.5Zm0 3.25a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5h-5.5Zm-.75 4a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 0 1.5h-5.5a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}