package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroceryOrSupermarket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<circle cx="44" cy="42" r="4" fill="currentColor"/><circle cx="15" cy="42" r="4" fill="currentColor"/><path fill="currentColor" d="M47 33H15.771l.667-1.082c.286-.464.37-1.025.233-1.553l-.651-2.506l28.983-1.506C46.102 26.297 47 25.35 47 24.25V11c0-1.1-.9-2-2-2H11.119l-.391-1.503A2 2 0 0 0 8.792 6H2a2 2 0 0 0 0 4h5.246l5.34 20.545l-2.1 3.405a1.998 1.998 0 0 0-.043 2.024A1.997 1.997 0 0 0 12.188 37H47a2 2 0 0 0 0-4z"/>`),
		g.Group(children),
	)
}