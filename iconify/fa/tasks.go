package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1024 1280h640v-128h-640v128zM640 768h1024V640H640v128zm640-512h384V128h-384v128zm512 832v256q0 26-19 45t-45 19H64q-26 0-45-19t-19-45v-256q0-26 19-45t45-19h1664q26 0 45 19t19 45zm0-512v256q0 26-19 45t-45 19H64q-26 0-45-19T0 832V576q0-26 19-45t45-19h1664q26 0 45 19t19 45zm0-512v256q0 26-19 45t-45 19H64q-26 0-45-19T0 320V64q0-26 19-45T64 0h1664q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}