package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Game(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M143 96V24h72v72h-72zm359-72h71v72h-71V24zm143 215V96h72v287h-72v71h-72v72h72v72h72v71H573v-71h-71v-72H215v72h-72v71H0v-71h72v-72h71v-72H72v-71H0V96h72v143h71v-71h72V96h72v72h143V96h72v72h71v71h72zm-358 72v-72h-72v72h72zm215 0v-72h-72v72h72z"/>`),
		g.Group(children),
	)
}