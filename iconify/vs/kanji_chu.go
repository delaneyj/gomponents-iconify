package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KanjiChu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="m530 1079l-226-75q-32 112-115 313q-87 212-178 344l225 131q87-140 174-351q90-218 120-362zm-67-330Q347 633 133 471L0 664q185 148 322 290zm83-435q-96-96-333-282L71 219q186 156 330 300zm1246 1442v-233h-503v-387h422V908h-422V596h471V371h-432l56-45Q1231 163 1049 0L850 151q148 140 222 220H553v225h469v312H585v228h437v387H480v233h1312z"/>`),
		g.Group(children),
	)
}