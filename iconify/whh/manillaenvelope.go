package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Manillaenvelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800.424 1024h-768q-13 0-22.5-9.5T.424 992V32q0-13 9.5-22.5t22.5-9.5h768q13 0 22.5 9.5t9.5 22.5v960q0 13-9.5 22.5t-22.5 9.5zm-736-64h320V570q-28-10-46-34.5t-18-55.5t18-55.5t46-34.5v-70q-12 0-130-26t-190-46v712zm704-896h-704v120q101 28 258 61q-2-11-2-21q0-40 28-68t68-28t68 28t28 68q0 10-2 21q157-33 258-61V64zm0 184q-72 20-190 46t-130 26v70q28 10 46 34.5t18 55.5t-18 55.5t-46 34.5v390h320V248z"/>`),
		g.Group(children),
	)
}