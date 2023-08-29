package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CannedChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1408H731l-475 475v-475H0V128zm1920 1280V256H128v1152h256v293l293-293h1243zm-256-896v128H384V512h1280zm0 256v128H384V768h1280zm0 256v128H384v-128h1280z"/>`),
		g.Group(children),
	)
}