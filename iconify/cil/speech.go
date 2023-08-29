package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speech(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M144 240h32v32h-32zm96 0h32v32h-32zm96 0h32v32h-32z"/><path fill="currentColor" d="M464 32H48a32.036 32.036 0 0 0-32 32v288a32.036 32.036 0 0 0 32 32h64v112h30.627l112-112H464a32.036 32.036 0 0 0 32-32V64a32.036 32.036 0 0 0-32-32Zm0 320H241.373L144 449.373V352H48V64h416Z"/>`),
		g.Group(children),
	)
}