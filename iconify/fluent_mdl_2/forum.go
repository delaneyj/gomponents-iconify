package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1408H731l-475 475v-475H0V128zm1920 1280V256H128v1152h256v293l293-293h1243zM1792 384v896H256V384h1536zM512 1152V896H384v256h128zm0-384V512H384v256h128zm1152 384V896H640v256h1024zm0-384V512H640v256h1024z"/>`),
		g.Group(children),
	)
}