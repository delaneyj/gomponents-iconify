package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemovePin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 14.5L3 21M7.676 7.89l-.979-.102L5 9.485l9.193 9.193l1.697-1.697l-.102-.981m-4.303-9l3.672-4.329l5.85 5.85l-4.308 3.654M3 3l18 18"/>`),
		g.Group(children),
	)
}