package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1152 704v128q0 221-147.5 384.5T640 1404v132h256q26 0 45 19t19 45t-19 45t-45 19H256q-26 0-45-19t-19-45t19-45t45-19h256v-132q-217-24-364.5-187.5T0 832V704q0-26 19-45t45-19t45 19t19 45v128q0 185 131.5 316.5T576 1280t316.5-131.5T1024 832V704q0-26 19-45t45-19t45 19t19 45zM896 320v512q0 132-94 226t-226 94t-226-94t-94-226V320q0-132 94-226T576 0t226 94t94 226z"/>`),
		g.Group(children),
	)
}