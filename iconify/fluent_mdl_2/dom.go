package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DOM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 1280h-512v-256H640v512h384v-128h512v512h-512v-256H512V512H384V0h512v512H640v384h384V768h512v512zm-398-398v284h284V882h-284zM498 398h284V114H498v284zm640 1124v284h284v-284h-284z"/>`),
		g.Group(children),
	)
}