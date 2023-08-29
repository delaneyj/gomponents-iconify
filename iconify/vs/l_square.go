package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm43 1419q0 31 22 52.5t53 21.5h884q31 0 53-21.5t22-52.5v-150q0-31-22-52.5t-53-21.5H678V373q0-31-22-52.5T603 299H454q-31 0-53 21.5T379 373v1046z"/>`),
		g.Group(children),
	)
}