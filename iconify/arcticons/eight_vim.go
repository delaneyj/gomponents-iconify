package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightVim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.43 22.145c1.877 1.757 1.636 2.969 1.585 4.657c-.254 8.404.705 11.65 2.398 12.47c5.307 1.596 14.768-6.516 8.601-22.906c-2.686-6.643-13.982-11.51-22.809-4.156c-2.729 2.58-3.8 5.26 0 9.665c2.531 2.665 3.621 4.784-1.256 4.64c-10.636-.7-7.176 5.41-5.799 7.054c5.558 6.435 12.11 7.362 15.27 3.48c15.497-21.227 22.717-15.587 22.036-11.985c-.777 7.624-8.32 11.543-13.795 12.154c-6.34.708-14.073-1.81-17.25-8.799c-2.262-5.724 8.405-7.473 14.147-8.816"/>`),
		g.Group(children),
	)
}