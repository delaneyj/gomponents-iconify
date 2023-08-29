package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1792 1216v128q0 26-19 45t-45 19H64q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h1664q26 0 45 19t19 45zm-384-384v128q0 26-19 45t-45 19H448q-26 0-45-19t-19-45V832q0-26 19-45t45-19h896q26 0 45 19t19 45zm256-384v128q0 26-19 45t-45 19H192q-26 0-45-19t-19-45V448q0-26 19-45t45-19h1408q26 0 45 19t19 45zM1280 64v128q0 26-19 45t-45 19H576q-26 0-45-19t-19-45V64q0-26 19-45t45-19h640q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}