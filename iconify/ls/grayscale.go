package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grayscale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 646V48h768v598H0zm383-114v63h64v-63h65v-61h-65v-63h65v-61h-65v-62h65v-61h-65v-63h65V99h-65v62h-64V99h-65v62h-62V99H51v496h267v-63h65zm0-371v63h-65v-63h65zM256 285v-61h62v61h-62zm191 0h-64v-61h64v61zm-129 0h65v62h-65v-62zm-62 123v-61h62v61h-62zm191 0h-64v-61h64v61zm-129 63v-63h65v63h-65zm-62 61v-61h62v61h-62zm127-61h64v61h-64v-61z"/>`),
		g.Group(children),
	)
}