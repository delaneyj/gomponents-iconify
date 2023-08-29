package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DamagedHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M87.195 53.838v79.494h44.213V53.838H87.195zm344.291 89.422c.34 7.22.677 14.441 1.014 21.662l27.861 41.004l-46.379 17.504l9.409 16.57l-24.334 32.486h86.273V143.26h-53.844zm-387.562 2.303v124.619H266.61l5.389-54.61l-63.18-17.166l21.7-38.656l-9.46-14.188H43.925zm6.709 134.802v201.711h53.316V321.408h96.614v160.668h271.152v-201.71h-83.766l-34.537 13.61l-23.178 30.768l-34.505-29.69l-26.827-14.689H50.632z"/>`),
		g.Group(children),
	)
}