package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChefHatBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M18.999 18H5.002c.01 1.397.082 2.912.584 3.414C6.172 22 7.115 22 9 22h6c1.886 0 2.829 0 3.415-.586c.502-.502.573-2.017.584-3.414Z"/><path d="M7 5a5 5 0 0 0-2 9.584V18h14v-3.416a5.001 5.001 0 0 0-2.737-9.53a4.502 4.502 0 0 0-8.526 0A5.04 5.04 0 0 0 7 5Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}