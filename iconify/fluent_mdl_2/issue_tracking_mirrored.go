package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTrackingMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 0h1536v2048H256V0zm1408 1920V128H384v1792h1280zM1024 512v128H512V512h512zm0 512v128H512v-128h512zm0 512v128H512v-128h512zm493-1107l-237 237l-173-173l90-90l83 83l147-147l90 90zm0 512l-237 237l-173-173l90-90l83 83l147-147l90 90zm0 512l-237 237l-173-173l90-90l83 83l147-147l90 90z"/>`),
		g.Group(children),
	)
}