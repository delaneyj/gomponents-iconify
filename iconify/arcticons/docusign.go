package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Docusign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.962 20.913h9.216L24 35.02L9.822 20.913h9.216V6.497h9.924v14.415ZM42.5 41.502h-37"/>`),
		g.Group(children),
	)
}