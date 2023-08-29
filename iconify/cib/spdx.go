package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spdx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 0v32h10.964l5.573-6.328v-8.073h8.766l6.698-6.635V0zm6.932 6.932h18.391L16.266 16c-4.979 4.984-9.115 9.068-9.193 9.068S6.932 20.985 6.932 16zm11.203 11.735V32H32V18.667z"/>`),
		g.Group(children),
	)
}