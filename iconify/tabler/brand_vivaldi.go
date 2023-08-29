package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandVivaldi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21.648 6.808c-2.468 4.28-4.937 8.56-7.408 12.836c-.397.777-1.366 1.301-2.24 1.356c-.962.102-1.7-.402-2.154-1.254c-1.563-2.684-3.106-5.374-4.66-8.064c-.943-1.633-1.891-3.266-2.83-4.905a2.47 2.47 0 0 1-.06-2.45A2.493 2.493 0 0 1 4.381 3.02a2.39 2.39 0 0 1 2.287 1.281c.697 1.19 2.043 3.83 2.55 4.682A3.919 3.919 0 0 0 12.5 11c2.126.133 3.974-.95 4.21-3.058c0-.164.228-3.178.846-3.962c.619-.784 1.64-1.155 2.606-.893a2.484 2.484 0 0 1 1.814 2.062a2.57 2.57 0 0 1-.343 1.674"/>`),
		g.Group(children),
	)
}