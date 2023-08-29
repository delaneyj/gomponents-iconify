package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hdr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 104v288h96a32.036 32.036 0 0 0 32-32V136a32.036 32.036 0 0 0-32-32Zm96 256h-64V136h64Zm-128 32V104h-32v128H64V104H32v288h32V264h64v128h32zm320-160v-96a32.036 32.036 0 0 0-32-32h-96v288h32V264h29.167L440 331.081V392h32v-67.081L447.632 264H448a32.036 32.036 0 0 0 32-32Zm-96 0v-96h64v96Z"/>`),
		g.Group(children),
	)
}