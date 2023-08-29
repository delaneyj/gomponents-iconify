package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H88a20 20 0 0 0-20 20v12H48a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h128a20 20 0 0 0 20-20v-20h12a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20ZM92 52h112v59.72l-9.86-9.86a20 20 0 0 0-28.28 0L103.72 164H92Zm80 152H52V84h16v84a20 20 0 0 0 20 20h84Zm-34.34-40L180 121.66l24 24V164ZM108 88a20 20 0 1 1 20 20a20 20 0 0 1-20-20Z"/>`),
		g.Group(children),
	)
}