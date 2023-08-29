package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RocketOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRocketOne0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="m20.906 6.063l1.43-.954a3 3 0 0 1 3.328 0l1.43.954A20 20 0 0 1 36 22.703V33H12V22.704a20 20 0 0 1 8.906-16.641Z"/><circle cx="24" cy="20" r="6" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="m12 22l-6 6.217V33h36v-4.783L36 22M24 38v6m-8-4v2m16-2v2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRocketOne0)"/>`),
		g.Group(children),
	)
}