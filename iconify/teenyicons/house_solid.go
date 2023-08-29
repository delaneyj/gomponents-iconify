package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.79.093a.5.5 0 0 0-.58 0l-7 5A.5.5 0 0 0 0 5.5v9a.5.5 0 0 0 .5.5H2V8h5v7h7.5a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.21-.407L14 4.528V2h-1v1.814L7.79.094ZM11 12V8h1v4h-1Z" clip-rule="evenodd"/><path fill="currentColor" d="M6 15v-3H5v-1h1V9H3v6h3Z"/>`),
		g.Group(children),
	)
}