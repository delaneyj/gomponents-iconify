package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SteelClaws(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M20.11 16.705h120.31l300.66 207.21l56.39 134l-138.88-96l-7.06-16.79zM309 423.295l-56.39-134l-238.08-164.09v94.45zm-48.47-146.43l10.79 25.64l128.76 89l-56.39-134l-329.16-226.8v76.64z"/>`),
		g.Group(children),
	)
}