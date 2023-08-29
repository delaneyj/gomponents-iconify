package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountBoxMail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m448 107l-64 42l-64-42V85l64 43l64-43v22zM469 0q18 0 30.5 12.5T512 43v298q0 18-12.5 30.5T469 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h426zM170.5 64q-26.5 0-45 18.5T107 128t18.5 45.5t45 18.5t45.5-18.5t19-45.5t-19-45.5T170.5 64zM299 320v-21q0-20-24-36t-52.5-23t-52-7t-52 7t-52 23T43 299v21h256zm170-128V64H299v128h170z"/>`),
		g.Group(children),
	)
}