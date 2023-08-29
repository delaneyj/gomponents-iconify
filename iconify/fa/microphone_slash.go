package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m271 945l-101 101q-42-103-42-214V704q0-26 19-45t45-19t45 19t19 45v128q0 53 15 113zm1114-602l-361 361v128q0 132-94 226t-226 94q-55 0-109-19l-96 96q97 51 205 51q185 0 316.5-131.5T1152 832V704q0-26 19-45t45-19t45 19t19 45v128q0 221-147.5 384.5T768 1404v132h256q26 0 45 19t19 45t-19 45t-45 19H384q-26 0-45-19t-19-45t19-45t45-19h256v-132q-125-13-235-81l-254 254q-10 10-23 10t-23-10l-82-82q-10-10-10-23t10-23L1257 215q10-10 23-10t23 10l82 82q10 10 10 23t-10 23zm-380-132L384 832V320q0-132 94-226T704 0q102 0 184.5 59T1005 211z"/>`),
		g.Group(children),
	)
}