package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DependencyRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 512v640h512v128H0V384h1280v512h-128V512H128zm640 1408v-896h1280v896H768zm128-768v640h1024v-640H896zM2048 92l-228 228l228 228l-91 92l-229-229l-228 229l-92-91l229-229l-229-229l92-91l228 229L1957 0l91 92z"/>`),
		g.Group(children),
	)
}