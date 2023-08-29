package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftOutlookLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 92a36 36 0 1 0 36 36a36 36 0 0 0-36-36Zm0 48a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm128-40h-4V40a20 20 0 0 0-20-20h-80a20 20 0 0 0-20 20v16H36a20 20 0 0 0-20 20v104a20 20 0 0 0 20 20h32v16a20 20 0 0 0 20 20h128a20 20 0 0 0 20-20v-96a20 20 0 0 0-20-20Zm-44.45 68L212 136.54v62.92ZM116 44h72v80.8l-28 21.78V76a20 20 0 0 0-20-20h-24ZM40 80h96v96H40Zm52 120h48a20 20 0 0 0 18.28-11.92L189 212H92Z"/>`),
		g.Group(children),
	)
}