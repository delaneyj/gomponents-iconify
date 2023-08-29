package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Saikou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" d="M18.277 42.5s14.171-8.083 14.295-13.28c.07-2.88-1.029-4.205-3.003-4.804M18.47 23.292c-2.066-.83-2.78-2.468-2.563-5.045c.337-4.013 12.127-10.264 14.002-12.747"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.242 24.75l-3.375 1.948l-5.973 3.449a.866.866 0 0 1-1.299-.75V18.603a.866.866 0 0 1 1.3-.75l3.374 1.949l5.973 3.448a.866.866 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}