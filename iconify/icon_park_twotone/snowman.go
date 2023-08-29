package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSnowman0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m12 24l-8-4m4 2v-4m32 4v-4"/><circle cx="24" cy="10" r="6" fill="#555" stroke="#fff" stroke-width="4"/><ellipse cx="24" cy="30" fill="#555" stroke="#fff" stroke-width="4" rx="13" ry="14"/><circle cx="24" cy="26" r="2" fill="#fff"/><circle cx="24" cy="31" r="2" fill="#fff"/><circle cx="24" cy="36" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m44 20l-8 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSnowman0)"/>`),
		g.Group(children),
	)
}