package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaitlistConfirmMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1344 998l147-147l90 90l-237 237l-173-173l90-90l83 83zm-832 538h512v128H512v-128zm512-896H512V512h512v128zm0 512H512v-128h512v128zm557-723l-237 237l-173-173l90-90l83 83l147-147l90 90zm-426 1491l128 128H256V0h1536v1283l-128 128V128H384v1792h771zm874-467l-557 558l-269-270l90-90l179 178l467-466l90 90z"/>`),
		g.Group(children),
	)
}