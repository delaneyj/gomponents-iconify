package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 640v1280H512v-512H128V128h1408v512h384zm-128 1152v-768H640v768h1152zM512 1280V640h896V512H256v768h256zM256 256v128h1152V256H256zm1536 512H640v128h1152V768z"/>`),
		g.Group(children),
	)
}