package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareviber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-576-640v-96q0-27-16.5-53t-36.5-37t-31-1l-73 74q-30 30-34 80.5t15.5 110t62.5 126t103 126.5t126.5 103t126 63t110 16t80.5-34l74-74q10-10-1-30.5t-37-37t-53-16.5h-96l-65 40q-45-12-111-65t-119-119t-65-111zm382 160l65 11q-8-120-93-205t-204-93l10 65q87 11 149 73t73 149zm-210-155l12 72q39 19 59 59l72 12q-13-53-51.5-91.5t-91.5-51.5zm-44-261l11 65q100 2 184.5 52.5t135 135t52.5 184.5l65 11q0-91-35.5-174t-95.5-143t-143-95.5t-174-35.5z"/>`),
		g.Group(children),
	)
}