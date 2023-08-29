package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbstractNinetyEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M115 21v94H21v94h94v94H21v94h94v94h94v-94h94v94h94v-94h94v-94h-94v-94h94v-94h-94V21h-94v94h-94V21h-94zm141 148.931c47.548 0 86.069 38.521 86.069 86.069s-38.52 86.069-86.069 86.069s-86.069-38.521-86.069-86.069s38.521-86.069 86.07-86.069zm0 29.963c-30.978 0-56.106 25.128-56.106 56.106s25.128 56.106 56.106 56.106s56.106-25.128 56.106-56.106s-25.128-56.106-56.106-56.106z"/>`),
		g.Group(children),
	)
}