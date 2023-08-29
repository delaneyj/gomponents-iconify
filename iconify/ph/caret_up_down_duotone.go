package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUpDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M80 176h96l-48 48Zm48-144L80 80h96Z" opacity=".2"/><path d="M176 168H80a8 8 0 0 0-5.66 13.66l48 48a8 8 0 0 0 11.32 0l48-48A8 8 0 0 0 176 168Zm-48 44.69L99.31 184h57.38ZM80 88h96a8 8 0 0 0 5.66-13.66l-48-48a8 8 0 0 0-11.32 0l-48 48A8 8 0 0 0 80 88Zm48-44.69L156.69 72H99.31Z"/></g>`),
		g.Group(children),
	)
}