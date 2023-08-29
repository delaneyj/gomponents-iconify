package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusCircleBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 256q115 0 221 29t199 84t168 130t130 168t84 199t30 222q0 115-29 221t-84 199t-130 168t-168 130t-199 84t-222 30q-115 0-221-29t-199-84t-168-130t-130-168t-84-199t-30-222q0-115 29-221t84-199t130-168t168-130t199-84t222-30zm0 1536q124 0 239-41t211-122L419 638q-80 95-121 210t-42 240q0 97 25 187t71 168t110 142t143 111t168 71t187 25zm541-254q80-95 121-210t42-240q0-97-25-187t-71-168t-110-142t-143-111t-168-71t-187-25q-124 0-239 41T510 547l991 991z"/>`),
		g.Group(children),
	)
}