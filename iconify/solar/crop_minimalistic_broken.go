package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 19h-9c-3.771 0-5.657 0-6.828-1.172c-.654-.653-.943-1.528-1.07-2.828M5 11V2m3 3h3c3.771 0 5.657 0 6.828 1.172c.654.653.943 1.528 1.07 2.828M2 5h3m14 14v3m0-9v3"/>`),
		g.Group(children),
	)
}