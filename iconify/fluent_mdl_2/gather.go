package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1920v-512h512v512H0zm128-384v256h256v-256H128zm1408-896h512v896h-512V640zm384 768v-256h-256v256h256zm0-384V768h-256v256h256zM0 1280V768h512v512H0zm128-384v256h256V896H128zM0 640V128h512v512H0zm128-384v256h256V256H128zm873 605l90-90l317 317l-317 317l-90-90l163-163H640v-128h524l-163-163z"/>`),
		g.Group(children),
	)
}