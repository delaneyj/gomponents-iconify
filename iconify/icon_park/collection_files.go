package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionFiles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M8 44V4H31L40 14.5V44H8Z"/><path fill="#43CCF8" stroke="#fff" d="M24 15L27.0841 21.7551L34.4616 22.6008L28.9902 27.6214L30.4656 34.8992L24 31.247L17.5344 34.8992L19.0098 27.6214L13.5384 22.6008L20.9159 21.7551L24 15Z"/></g>`),
		g.Group(children),
	)
}