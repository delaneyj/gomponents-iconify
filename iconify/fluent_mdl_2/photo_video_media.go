package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoVideoMedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 512h128v128H256V512zm0 256h128v128H256V768zm1024-128h-128V512h128v128zm768 128v1152H512v-640H0V128h1536v640h512zm-128 128H640v549l320-319l448 447l320-319l192 191V896zM512 1152V768h896V256h-128v128h-128V256H384v128H256V256H128v896h128v-128h128v128h128zm128 640h805l-485-486l-320 321v165zm1280 0v-165l-192-193l-229 230l128 128h293zm-448-640q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}