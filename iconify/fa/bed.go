package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M256 768h1728q26 0 45 19t19 45v448h-256v-256H256v256H0V64q0-26 19-45T64 0h128q26 0 45 19t19 45v704zm576-320q0-106-75-181t-181-75t-181 75t-75 181t75 181t181 75t181-75t75-181zm1216 256v-64q0-159-112.5-271.5T1664 256H960q-26 0-45 19t-19 45v384h1152z"/>`),
		g.Group(children),
	)
}