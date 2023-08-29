package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1216q0 115-30 221t-84 198t-130 169t-168 130t-199 84t-221 30q-115 0-221-30t-198-84t-169-130t-130-168t-84-199t-30-221v-64h128v64q0 97 25 187t71 168t110 142t143 111t168 71t187 25q97 0 187-25t168-71t142-110t111-143t71-168t25-187q0-97-25-187t-71-168t-110-142t-143-111t-168-71t-187-25H250l291 291l-90 90L6 448L451 3l90 90l-291 291h838q115 0 221 30t198 84t169 130t130 168t84 199t30 221z"/>`),
		g.Group(children),
	)
}