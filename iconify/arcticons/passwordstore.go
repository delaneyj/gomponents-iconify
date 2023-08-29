package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passwordstore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35 25.677a3 3 0 1 1-3-3.002a3.013 3.013 0 0 1 3 3.002Zm-8.354 7.498a3 3 0 1 1-3-3.002a3.013 3.013 0 0 1 3 3.002ZM35 37.077a3 3 0 1 1-3-3.002a3.013 3.013 0 0 1 3 3.002Zm-3-14.402V12.286A7.786 7.786 0 0 0 24.214 4.5h-.428A7.786 7.786 0 0 0 16 12.286V17.5m16 16.575v-5.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.521 17.5h3.317a4 4 0 0 1 3.953 4.608l-2.77 18a4 4 0 0 1-3.953 3.392H12.932a4 4 0 0 1-3.954-3.392l-2.769-18a4 4 0 0 1 3.953-4.608h19.36m-3.65 13.675l3.903-3.503"/>`),
		g.Group(children),
	)
}