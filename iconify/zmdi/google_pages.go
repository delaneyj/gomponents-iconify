package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 43q0-18 12.5-30.5T43 0h128v107L85 85l22 86H0V43zm107 170l-22 86l86-22v107H43q-18 0-30.5-12.5T0 341V213h107zm192 86l-22-86h107v128q0 18-12.5 30.5T341 384H213V277zM341 0q18 0 30.5 12.5T384 43v128H277l22-86l-86 22V0h128z"/>`),
		g.Group(children),
	)
}