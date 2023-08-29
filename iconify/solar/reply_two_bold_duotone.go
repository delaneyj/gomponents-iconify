package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10.03 17.53a.75.75 0 0 1-1.06 0l-5-5a.75.75 0 0 1 0-1.06l5-5a.75.75 0 1 1 1.06 1.06L5.56 12l4.47 4.47a.75.75 0 0 1 0 1.06Z" clip-rule="evenodd"/><path d="M6.31 12.75h8.19c.953 0 2.367-.28 3.563-1.141c1.235-.89 2.187-2.365 2.187-4.609a.75.75 0 0 0-1.5 0c0 1.756-.715 2.78-1.563 3.391c-.887.639-1.974.859-2.687.859H6.31l-.75.75l.75.75Zm-2.503-.463Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}