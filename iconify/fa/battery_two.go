package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M256 1024V256h896v768H256zm1920-704q53 0 90.5 37.5T2304 448v384q0 53-37.5 90.5T2176 960v160q0 66-47 113t-113 47H160q-66 0-113-47T0 1120V160Q0 94 47 47T160 0h1856q66 0 113 47t47 113v160zm0 512V448h-128V160q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v960q0 14 9 23t23 9h1856q14 0 23-9t9-23V832h128z"/>`),
		g.Group(children),
	)
}