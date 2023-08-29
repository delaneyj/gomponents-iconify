package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Livemagicscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.453 23.63c.564 3.023.226 3.331-.769 6.465m-3.322 7.592A19.294 19.294 0 1 1 14.92 7.853m7-1.33a19.319 19.319 0 0 1 5.126.688m-10.511-.266l1.8-1.35m10.645-.763l2.001.175m-1.707 5.823l.788 1.238m1.9-3.505l1.717.217M38.9 35.464l1.8-1.5m-2.325-9.232l-.45-2.1m5.927-.976c1.963-1.473-.416.416.9-.9a1.068 1.068 0 0 1 .301-.15m-3.609-1.552l-.287-1.325m-3.287.646l-.943-1.092M20.694 3.537L22.934 3m16.727 8.416l.969-.044l.138-.118m-1.886 2.996l1.086.552m-2.863-6.798l1.187.286m5.29 8.21l.509-1.257"/>`),
		g.Group(children),
	)
}