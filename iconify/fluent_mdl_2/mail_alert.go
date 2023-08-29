package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v684q-55-79-128-139V648l-282 140q-41-9-82-14t-84-6q-23 0-46 2t-47 5l526-263H143l864 432q-25 23-48 47t-43 50L128 648v888h643q3 32 9 64t15 64H0V384h2048zm-576 512q119 0 224 45t183 124t123 183t46 224q0 119-45 224t-124 183t-183 123t-224 46q-119 0-224-45t-183-124t-123-183t-46-224q0-119 45-224t124-183t183-123t224-46zm0 1024q93 0 174-35t142-96t96-142t36-175q0-93-35-174t-96-142t-142-96t-175-36q-93 0-174 35t-142 96t-96 142t-36 175q0 93 35 174t96 142t142 96t175 36zm-64-768h128v384h-128v-384zm0 512h128v128h-128v-128z"/>`),
		g.Group(children),
	)
}