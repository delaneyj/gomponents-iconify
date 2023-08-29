package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulkUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M512 1920h1024v128H384v-256H128V0h859l402 402l403 403v91h-640V384H512v1536zm768-1152h293l-293-293v293zm-896 896V256h677L933 128H256v1536h128zm1344-410l317 317l-90 90l-163-162v549h-128v-549l-163 162l-90-90l317-317zm320-230v128h-640v-128h640z"/>`),
		g.Group(children),
	)
}