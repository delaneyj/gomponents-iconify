package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sliders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M352 1152v128H0v-128h352zm352-128q26 0 45 19t19 45v256q0 26-19 45t-45 19H448q-26 0-45-19t-19-45v-256q0-26 19-45t45-19h256zm160-384v128H0V640h864zM224 128v128H0V128h224zm1312 1024v128H800v-128h736zM576 0q26 0 45 19t19 45v256q0 26-19 45t-45 19H320q-26 0-45-19t-19-45V64q0-26 19-45t45-19h256zm640 512q26 0 45 19t19 45v256q0 26-19 45t-45 19H960q-26 0-45-19t-19-45V576q0-26 19-45t45-19h256zm320 128v128h-224V640h224zm0-512v128H672V128h864z"/>`),
		g.Group(children),
	)
}