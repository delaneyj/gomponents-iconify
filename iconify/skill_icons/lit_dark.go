package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LitDark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#242938" rx="60"/><path fill="#00E8FF" d="m88 148l20-60l90 90l-30 50l-40-40h-20"/><path fill="#283198" d="M128 188v-80l40-40v80M48 188l40 40l20-40l-20-40H68"/><path fill="#324FFF" d="M88 148V68l40-40v80m40 120v-80l40-40v80m-160 0v-80l40 40"/><path fill="#0FF" d="M88 228v-80l40 40"/></g>`),
		g.Group(children),
	)
}