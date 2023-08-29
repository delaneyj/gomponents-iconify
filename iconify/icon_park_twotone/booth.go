package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Booth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBooth0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 5h40v8l-1.398.84a7 7 0 0 1-7.203 0L34 13l-1.398.84a7 7 0 0 1-7.203 0L24 13l-1.398.84a7 7 0 0 1-7.204 0L14 13l-1.399.84a7 7 0 0 1-7.202 0L4 13V5Z"/><path d="M6 25h36v18H6zm3-9v9m30-9v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBooth0)"/>`),
		g.Group(children),
	)
}