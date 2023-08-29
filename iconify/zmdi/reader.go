package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 192h149v32H256v-32zm0-21h149h-149zm0 106h149h-149zM427 21q17 0 29.5 12.5T469 64v277q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V64q0-18 12.5-30.5T43 21h384zm0 320V64H235v277h192z"/>`),
		g.Group(children),
	)
}