package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Documentation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 512v1408H768v-256H512v-256H256V0h731l256 256h421v256h256zm-896-128h165l-165-165v165zm256 896V512H896V128H384v1152h896zm256 256V384h-128v1024H640v128h896zm257-896h-129v1024H896v128h897V640z"/>`),
		g.Group(children),
	)
}