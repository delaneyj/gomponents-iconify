package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1664 448v256q0 26-19 45t-45 19h-64q-26 0-45-19t-19-45V448q0-106-75-181t-181-75t-181 75t-75 181v192h96q40 0 68 28t28 68v576q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V736q0-40 28-68t68-28h672V448q0-185 131.5-316.5T1216 0t316.5 131.5T1664 448z"/>`),
		g.Group(children),
	)
}