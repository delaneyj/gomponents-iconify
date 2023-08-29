package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBabyMobile0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4.5 18L24 8l20 10m-20 0V4M10 30V16"/><circle cx="10" cy="36" r="6" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M38 30V16"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="m32 36l6-6l6 6l-6 6l-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 23v-5H19v5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBabyMobile0)"/>`),
		g.Group(children),
	)
}