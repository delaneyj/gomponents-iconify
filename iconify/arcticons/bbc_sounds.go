package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BbcSounds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.69 6.059H42.5v35.883H23.69zm-11.299 5.587H23.69v24.708H12.391zM5.5 16.55h6.891v14.899H5.5z"/>`),
		g.Group(children),
	)
}