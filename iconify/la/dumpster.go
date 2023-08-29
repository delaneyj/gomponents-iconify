package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dumpster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m5 6l-2 8h1.334l.332 2H3v2h2l1 6v2h2v-2h16v2h2v-2l1-6h2v-2h-1.666l.332-2H29l-2-8H5zm1.563 2h2.775l-.766 4h-3.01l1-4zm4.828 0H15v4h-4.37l.76-4zM17 8h3.61l.76 4H17V8zm5.662 0h2.776l1 4h-3.01l-.766-4zm-16.3 6h19.277l-1.332 8H7.693l-1.332-8z"/>`),
		g.Group(children),
	)
}