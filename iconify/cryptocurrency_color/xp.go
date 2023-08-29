package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#008200"/><path fill="#fff" fill-rule="nonzero" d="M14.79 23h-1.798l-1.082-4.25l2.645-1.382l.861 2.46l.641-3.245l3.515-1.837l-.194.976l2.03.019c.784-.012 1.391-.277 1.823-.793c.432-.517.637-1.21.615-2.083a2.31 2.31 0 0 0-.035-.335l2.73-1.427c.378.695.522 1.514.434 2.455c-.146 1.372-.74 2.47-1.779 3.293c-1.039.824-2.364 1.236-3.973 1.236l-2.306-.01L17.94 23zm.421-10.99h1.75l-2.697 3.825l-6.66 3.481l2.468-3.49L7.7 9h3.48l1.23 4.789l1.052-1.78h1.75h-1.75L15.34 9h2.215l5.126.01c1.165.032 2.12.341 2.866.927l-2.8 1.464a1.773 1.773 0 0 0-.208-.026l-2.294-.02l-.3 1.51l-3.517 1.838l.532-2.694zm-9.425 9.942l5.327-2.785L8.765 23H5l.767-1.085z"/></g>`),
		g.Group(children),
	)
}