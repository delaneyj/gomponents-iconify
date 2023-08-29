package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PicoEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M204.287 0v101.883l-102.142.262l-.262 102.142H0v103.426h101.889V410.11h102.398V512h103.426V410.111H410.11V307.713H512V204.287H410.111V101.89H307.713V0H204.287zm1.026 102.912h101.375v102.4h102.4v101.372l-102.145.26l-.26 102.144h-101.37v-102.4h-102.4V205.313h102.4v-102.4z"/>`),
		g.Group(children),
	)
}