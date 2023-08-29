package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhysicsWallah(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.624 25.54v-15.2h4.196a4.357 4.357 0 1 1 0 8.714h-4.196m8.133 5.558l3.227 12.91l3.227-12.91l3.228 12.91l3.227-12.91"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.036" d="M2.88 18.706A21.866 21.866 0 0 1 34.932 5.43M13.067 42.57a21.866 21.866 0 0 0 32.054-13.276"/>`),
		g.Group(children),
	)
}