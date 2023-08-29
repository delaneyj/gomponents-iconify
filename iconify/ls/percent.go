package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M62 819L596 0h77L137 819H62zM174 15c95 0 173 78 173 173c0 96-78 174-173 174C78 362 0 284 0 188C0 93 78 15 174 15zm0 285c61 0 111-51 111-112S235 78 174 78S63 127 63 188s50 112 111 112zm402 142c96 0 174 78 174 174c0 95-78 174-174 174c-95 0-173-79-173-174c0-96 78-174 173-174zm0 285c61 0 111-50 111-111s-50-111-111-111s-111 50-111 111s50 111 111 111z"/>`),
		g.Group(children),
	)
}