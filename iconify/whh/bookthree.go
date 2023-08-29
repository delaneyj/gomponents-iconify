package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookthree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M997.295 826q-28 12-185 72t-209 89q-8 4-28 12q-3 11-11.5 18.5t-19.5 7.5h-64q-11 0-19.5-7.5t-11.5-18.5q-20-8-28-12q-52-29-209.5-89t-184.5-72q-11-5-19-17.5t-8-25.5V88q0-14 8-20t19-1q51 22 115 58.5t97.5 58t117.5 76.5q30 20 91 32v-3q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5v3q61-12 91-32q84-55 117.5-76.5t97.5-58t115-58.5q11-5 19 1t8 20v695q0 13-8 25.5t-19 17.5zm-389-633h-6q136-120 271-190q10-5 16.5 1t6.5 20v28q-131 61-262 155q-10-14-26-14zm-218 14q-131-94-262-155V24q0-14 6.5-20t16.5-1q135 70 271 190h-6q-16 0-26 14z"/>`),
		g.Group(children),
	)
}