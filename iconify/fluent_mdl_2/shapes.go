package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shapes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1280h-419l370 640H595l267-463q-66 38-138 58t-148 21q-119 0-224-45t-183-124t-123-183T0 960q0-119 45-224t124-183t183-123t224-46q85 0 166 24t154 73V128h1152v1152zM576 1408q93 0 174-35t142-97t96-142t36-174q0-92-35-173t-96-143t-142-96t-175-36q-93 0-174 35t-142 96t-96 142t-36 175q0 93 35 174t96 142t142 96t175 36zm1201 384q-121-209-240-416t-240-416q-121 209-240 416t-240 416h960zm143-1536h-896v347q61 75 94 165t34 188l145-252l258 448h365V256z"/>`),
		g.Group(children),
	)
}