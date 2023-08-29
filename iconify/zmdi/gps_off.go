package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M425 219h44v42h-43q-4 36-21 68l-32-32q11-28 11-57q0-62-43.5-105.5T235 91q-30 0-57 11l-32-32q32-17 67-21V5h43v44q67 8 114.5 55.5T425 219zM43 75l27-27l357 357l-27 27l-44-44q-44 36-100 43v44h-43v-44q-66-8-114-55.5T44 261H0v-42h44q6-56 42-100zm283 283L116 149q-31 40-31 91q0 62 44 105.5T235 389q50 0 91-31z"/>`),
		g.Group(children),
	)
}