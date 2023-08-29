package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdhesivePlasterTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m20.91 13.332l-7.578 7.577a5.41 5.41 0 0 0 7.577-7.577Z"/><path fill-rule="evenodd" d="m12.235 19.885l7.65-7.65l-8.12-8.12l-7.65 7.65l8.12 8.12Zm-2.356-8.592a1 1 0 1 1-1.414 1.414a1 1 0 0 1 1.414-1.414Zm2.828 4.243a1 1 0 1 0-1.414-1.415a1 1 0 0 0 1.414 1.415Zm0-7.071a1 1 0 1 1-1.414 1.414a1 1 0 0 1 1.414-1.414Zm2.828 4.242a1 1 0 1 0-1.414-1.414a1 1 0 0 0 1.414 1.414Z" clip-rule="evenodd"/><path d="m3.09 10.668l7.578-7.577a5.41 5.41 0 0 0-7.577 7.577Z"/></g>`),
		g.Group(children),
	)
}