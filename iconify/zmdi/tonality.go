package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tonality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zM192 385V47q-64 8-106.5 56T43 216t43 113t106 56zm43-338v20h61q-29-16-61-20zm0 62v22h126q-7-12-15-22H235zm0 64v22h148q-2-9-5-22H235zm0 212q32-4 61-20h-61v20zm111-62q8-10 15-22H235v22h111zm32-64q3-13 5-22H235v22h143z"/>`),
		g.Group(children),
	)
}