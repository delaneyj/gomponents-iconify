package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tapscanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M40.646 24c.26.09.26.12.26 2.93v2.829l-.71.68a2.36 2.36 0 0 0-.8 1.15c-.12.45-.83 1.17-5 5l-5.66 5.23a6.638 6.638 0 0 1-1.94 1.47a5.21 5.21 0 0 1-1.599.21c-1.12.01-1.34-.03-9.07-1.65l-4-.84c-3.71-.75-3.94-.83-4.09-1.42c-.06-.24-.24-.35-.8-.48c-1.26-.3-1.32-.43-1.44-3.12a17.704 17.704 0 0 1 0-2.65c.07-.23 2.64-2.18 7.88-6c4.28-3.11 8.98-6.57 9.16-6.67m0 0c-.16-.14-2.5-3.38-4.68-6.27c-5.53-7.419-6.17-8.239-6.17-8.599s.96-1.2 1.48-1.3h0c1.276.038 2.549.155 3.81.35l7.11.92c6.08.78 6.53.86 7.24 1.22c.94.48 1.41 1.01 4.07 4.82l4.36 6.23c1.1 1.53 2.01 2.91 2.1 3.07a2.35 2.35 0 0 1-1.51 2.89"/><path d="M40.846 24.17L26.177 37.418a.94.94 0 0 1-.85.23l-19.53-4.36m17.04-12.619l17.809 3.37"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.009 25.943l-10.172 8.646a.553.553 0 0 1-.5.135l-12.35-2.64l11.34-8.411l11.682 2.27Z"/>`),
		g.Group(children),
	)
}