package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProtectedDocument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1280v768h-896v-768h128q0-76 17-145t56-123t99-84t148-32q87 0 147 31t99 85t56 122t18 146h128zm-640 0h384v-64q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75v64zm512 128h-640v512h640v-512zM384 1920h512v128H256V293L549 0h1243v860q-29-26-61-47t-67-37V128H603L384 347v1573z"/>`),
		g.Group(children),
	)
}