package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tiled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTiled0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTiled1" fill="currentColor" fill-rule="nonzero"><path id="feTiled2" d="M6 4h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm7 9v5h5v-5h-5Zm-7 0v5h5v-5H6Zm7-7v5h5V6h-5ZM6 6v5h5V6H6Z"/></g></g>`),
		g.Group(children),
	)
}