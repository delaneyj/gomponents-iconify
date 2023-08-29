package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Debloater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.89 18.944L14.68 43.5h20.092l2.79-24.556Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="4.75" d="M37.562 18.944L16.987 4.5"/><path fill="currentColor" d="M14.253 13.747a.75.75 0 1 1 1.06 0a.75.75 0 0 1-1.06 0Zm3.822 3.822a.75.75 0 1 1 1.06 0a.75.75 0 0 1-1.06 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.516 18.944a6.622 6.622 0 0 0-1.44-7.206h0a6.62 6.62 0 0 0-9.362 0h0a.945.945 0 0 0 0 1.336l5.87 5.87m.15-11.948l-.236 2.895m8.32 5.189l-2.895.236"/>`),
		g.Group(children),
	)
}