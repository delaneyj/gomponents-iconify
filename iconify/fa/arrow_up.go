package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1579 779q0 51-37 90l-75 75q-38 38-91 38q-54 0-90-38L992 651v704q0 52-37.5 84.5T864 1472H736q-53 0-90.5-32.5T608 1355V651L314 944q-36 38-90 38t-90-38l-75-75q-38-38-38-90q0-53 38-91L710 37q35-37 90-37q54 0 91 37l651 651q37 39 37 91z"/>`),
		g.Group(children),
	)
}