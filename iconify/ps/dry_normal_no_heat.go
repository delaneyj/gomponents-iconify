package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DryNormalNoHeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v426q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V43q0-18-12.5-30.5T469 0zm0 469H43V43h426v426zm-213-23q79 0 135.5-56T448 254t-56-136t-136-56t-136 56t-56 136t56 136t136 56z"/>`),
		g.Group(children),
	)
}