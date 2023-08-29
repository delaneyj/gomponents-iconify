package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1472 896q119 0 224 45t183 124t123 183t46 224q0 119-45 224t-124 183t-183 123t-224 46q-119 0-224-45t-183-124t-123-183t-46-224q0-119 45-224t124-183t183-123t224-46zm0 1024q93 0 174-35t142-96t96-142t36-175q0-93-35-174t-96-142t-142-96t-175-36q-93 0-174 35t-142 96t-96 142t-36 175q0 93 35 174t96 142t142 96t175 36zM640 768v512h122q-8 32-14 64t-9 64H512V640h293l384-384h91v506q-67 18-128 48V475L858 768H640zm1069 557l-147 147l147 147l-90 90l-147-146l-147 146l-90-90l146-147l-146-147l90-90l147 147l147-147l90 90z"/>`),
		g.Group(children),
	)
}