package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fullscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 235v64h64v42H0V235h43zM0 149V43h107v42H43v64H0zm256 150v-64h43v106H192v-42h64zM192 43h107v106h-43V85h-64V43z"/>`),
		g.Group(children),
	)
}