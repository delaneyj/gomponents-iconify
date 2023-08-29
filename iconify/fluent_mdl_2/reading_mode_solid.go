package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReadingModeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 1664q43 0 75 9t60 26t53 41t54 52H0V256h128v1408h512zm0-1536q67 0 132 16t124 50v1435q-54-45-120-69t-136-24H256V128h384zm1280 128v1536h-882q28-28 53-52t53-40t60-26t76-10h512V256h128zm-640 1280q-70 0-136 24t-120 69V194q59-33 124-49t132-17h384v1408h-384z"/>`),
		g.Group(children),
	)
}