package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M898 342v102q0 14-9 23t-23 9H698q-23 144-129 234T293 820q167 178 459 536q14 16 4 34q-8 18-29 18H532q-16 0-25-12Q201 1029 9 825q-9-9-9-22V676q0-13 9.5-22.5T32 644h112q132 0 212.5-43T459 476H32q-14 0-23-9t-9-23V342q0-14 9-23t23-9h413q-57-113-268-113H32q-13 0-22.5-9.5T0 165V32Q0 18 9 9t23-9h832q14 0 23 9t9 23v102q0 14-9 23t-23 9H631q47 61 64 144h171q14 0 23 9t9 23z"/>`),
		g.Group(children),
	)
}