package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vflat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.903 7.763l23.074 6.527a.718.718 0 0 1 .36 1.145l-9.814 11.968a1.847 1.847 0 0 1-2.154.528L5.1 16.706a.99.99 0 0 1-.094-1.775L17.294 8.07a3.434 3.434 0 0 1 2.609-.306Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.42 19.406a64.516 64.516 0 0 0-6.593 5.537a.689.689 0 0 0 .167 1.08C8.53 27.947 24.073 36.378 31.452 40.12a2.272 2.272 0 0 0 3.15-1.205c2.062-5.355 7.04-11.564 8.71-13.335a.54.54 0 0 0-.158-.859a37.238 37.238 0 0 0-5.715-2.093"/>`),
		g.Group(children),
	)
}