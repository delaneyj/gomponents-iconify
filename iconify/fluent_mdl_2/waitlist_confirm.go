package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaitlistConfirm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 640h-512V512h512v128zm-512 384h512v128h-512v-128zm131 896l128 128H256V0h1536v1283l-128 128V128H384v1792h771zM941 429L704 666L531 493l90-90l83 83l147-147l90 90zm0 512l-237 237l-173-173l90-90l83 83l147-147l90 90zm-237 569l147-147l90 90l-237 237l-173-173l90-90l83 83zm1325-57l-557 558l-269-270l90-90l179 178l467-466l90 90z"/>`),
		g.Group(children),
	)
}