package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Community(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.402 30.896h4m-4-6.9l2-1.1m0 0v8M7.032 5.876h33.936a1.519 1.519 0 0 1 1.516 1.521v24.79a1.519 1.519 0 0 1-1.516 1.521H5.515V7.398a1.519 1.519 0 0 1 1.517-1.522Zm9.028 27.847v8.4L5.516 33.709"/>`),
		g.Group(children),
	)
}