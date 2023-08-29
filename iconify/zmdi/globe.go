package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zM192 385v-41q-18 0-30.5-12.5T149 301v-21L47 178q-4 20-4 38q0 65 42.5 113T192 385zm147-54q45-49 45-115q0-53-29.5-96T277 58v9q0 17-12.5 29.5T235 109h-43v43q0 9-6.5 15t-14.5 6h-43v43h128q9 0 15 6.5t6 14.5v64h22q14 0 25 8.5t15 21.5z"/>`),
		g.Group(children),
	)
}