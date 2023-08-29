package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filmorago(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.78 14.26l9.74 9.762L24 43.5l-9.74-9.761z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.359 26.656l-.26-5.584l-6.818-6.833L24.042 4.5l9.739 9.761l-12.422 12.395z"/>`),
		g.Group(children),
	)
}