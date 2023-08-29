package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmplicationIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#7950ED" d="M0 130.519c0 67.317 49.315 122.613 112.77 130.34V.178C49.314 7.904 0 63.2 0 130.519Zm256-.006C256 62.853 206.18 7.213 142.053 0v261.199h113.78v-127.08c.167-1.201.167-2.403.167-3.606Z"/>`),
		g.Group(children),
	)
}