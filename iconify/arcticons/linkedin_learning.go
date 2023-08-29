package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinLearning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.5 9.5h-35c-1.1 0-2 .9-2 2v25c0 1.1.9 2 2 2h35c1.1 0 2-.9 2-2v-25c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.686 22.931l-9.145-5.265a1.233 1.233 0 0 0-1.848 1.069v10.53a1.233 1.233 0 0 0 1.848 1.069l9.145-5.265a1.233 1.233 0 0 0 0-2.137ZM24 9.5v10.006m0 8.836V38.5M8.25 31.991h12M8.25 24h10.442M8.25 16.009h12m7.5 15.982h12M30.303 24h9.447m-12-7.991h12"/>`),
		g.Group(children),
	)
}