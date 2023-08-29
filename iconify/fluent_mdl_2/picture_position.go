package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PicturePosition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M512 512h1024v1024H512V512zm128 896h549l-357-358l-192 193v165zm0-768v421l192-191l539 538h37V640H640zm576 256q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19zM859 347l-91-91L1024 0l256 256l-91 91l-165-166l-165 166zm330 1354l91 91l-256 256l-256-256l91-91l165 166l165-166zM347 859l-166 165l166 165l-91 91L0 1024l256-256l91 91zm1701 165l-256 256l-91-91l166-165l-166-165l91-91l256 256z"/>`),
		g.Group(children),
	)
}