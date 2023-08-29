package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 320v-43h86v43h-86zM0 64h384v43H0V64zm64 149v-42h256v42H64z"/>`),
		g.Group(children),
	)
}