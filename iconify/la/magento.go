package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magento(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3.227L5 10.174V22.44l3 1.713V11.826l8-5.053l8 5.053V24.15l3-1.712V10.174L16 3.227zm-2 8.351l-3 1.887v12.404l5 2.86l5-2.86V13.484L18 11.6v12.53l-2 1.141l-2-1.14V11.578z"/>`),
		g.Group(children),
	)
}