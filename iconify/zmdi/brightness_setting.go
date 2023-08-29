package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h384zm0 342V42H43v300h384zm-278-65v-53l-32-32l32-32v-53h54l32-32l32 32h53v53l32 32l-32 32v53h-53l-32 32l-32-32h-54zm86-149v128q26 0 45-18.5t19-45.5t-19-45.5t-45-18.5z"/>`),
		g.Group(children),
	)
}