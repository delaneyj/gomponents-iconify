package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneHelix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M12 9s-1.988-1.975-2-4c.001-1.993-.05-4.001 2-4c1.948.001 1.997 1.976 2 4c.003 1.985-2 4-2 4Zm3 3s1.975-1.988 4-2c1.993.001 4.001-.05 4 2c-.001 1.948-1.976 1.997-4 2c-1.985.003-4-2-4-2Zm-6 0s-1.975 1.988-4 2c-1.993-.001-4.001.05-4-2c.001-1.948 1.976-1.997 4-2c1.985-.003 4 2 4 2Zm3 3s1.988 1.975 2 4c-.001 1.993.05 4.001-2 4c-1.948-.001-1.997-1.976-2-4c-.003-1.985 2-4 2-4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}