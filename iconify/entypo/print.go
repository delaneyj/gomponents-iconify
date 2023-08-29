package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.501 6h17c.57 0 .477-.608.193-.707C18.409 5.194 15.251 4 14.7 4H14V1H6v3h-.699c-.55 0-3.709 1.194-3.993 1.293c-.284.099-.377.707.193.707zM19 7H1c-.55 0-1 .45-1 1v5c0 .551.45 1 1 1h2.283l-.882 5H17.6l-.883-5H19c.551 0 1-.449 1-1V8c0-.55-.449-1-1-1zM4.603 17l1.198-7.003H14.2L15.399 17H4.603z"/>`),
		g.Group(children),
	)
}