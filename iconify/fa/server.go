package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M128 1280h1024v-128H128v128zm0-512h1024V640H128v128zm1568 448q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68zM128 256h1024V128H128v128zm1568 448q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68zm0-512q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68zm96 832v384H0v-384h1792zm0-512v384H0V512h1792zm0-512v384H0V0h1792z"/>`),
		g.Group(children),
	)
}