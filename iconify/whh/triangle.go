package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 1024h-32v-32q0-13-9.5-22.5T928 960t-22.5 9.5T896 992v32h-64v-96q0-13-9.5-22.5T800 896t-22.5 9.5T768 928v96h-64v-32q0-13-9.5-22.5T672 960t-22.5 9.5T640 992v32h-64v-96q0-13-9.5-22.5T544 896t-22.5 9.5T512 928v96h-64v-32q0-13-9.5-22.5T416 960t-22.5 9.5T384 992v32h-64v-96q0-13-9.5-22.5T288 896t-22.5 9.5T256 928v96h-64v-32q0-13-9.5-22.5T160 960t-22.5 9.5T128 992v32H32q-13 0-22.5-9.5T0 992L992 0q13 0 22.5 9.5T1024 32v960q0 13-9.5 22.5T992 1024zM832 480q0-13-9.5-22.5T800 448L448 800q0 13 9.5 22.5T480 832h320q13 0 22.5-9.5T832 800V480z"/>`),
		g.Group(children),
	)
}