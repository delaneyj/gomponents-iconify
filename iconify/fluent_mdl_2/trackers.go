package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trackers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 256v1792H256V256h512q0-53 20-99t55-82t81-55t100-20q53 0 99 20t82 55t55 81t20 100h512zM640 512h768V384h-256V256q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50v128H640v128zm1024-128h-128v256H512V384H384v1536h1280V384zm-640 512h512v128h-512V896zm0 384h512v128h-512v-128zm0 384h512v128h-512v-128zM851 723l90 90l-237 237l-173-173l90-90l83 83l147-147zm0 384l90 90l-237 237l-173-173l90-90l83 83l147-147zm0 384l90 90l-237 237l-173-173l90-90l83 83l147-147z"/>`),
		g.Group(children),
	)
}