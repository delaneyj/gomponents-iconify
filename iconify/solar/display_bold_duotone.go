package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplayBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2.879 3.844C2 4.687 2 6.044 2 8.76v.96c0 2.715 0 4.073.879 4.916c.878.844 2.293.844 5.121.844h8c2.828 0 4.243 0 5.121-.844c.879-.843.879-2.2.879-4.916v-.96c0-2.715 0-4.073-.879-4.916C20.243 3 18.828 3 16 3H8c-2.828 0-4.243 0-5.121.844Z" clip-rule="evenodd"/><path d="M18.237 19.596L12.75 17.84v-2.36h-1.5v2.36l-5.487 1.756a.714.714 0 0 0-.474.911a.757.757 0 0 0 .948.455L12 19.118l5.763 1.845a.757.757 0 0 0 .949-.456a.714.714 0 0 0-.475-.91Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}