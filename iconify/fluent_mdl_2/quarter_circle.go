package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuarterCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v1792H128q0-99 6-192t23-183t42-180t65-182q68-160 164-300t215-258t258-209t294-157t320-97t341-34h64zm-128 129q-208 8-400 67t-360 161t-307 240t-240 307t-160 360t-68 400h1535V257z"/>`),
		g.Group(children),
	)
}