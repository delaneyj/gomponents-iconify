package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Progress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#5CE500" d="M63.017 260.487L2.769 223.492l60.248-35.871v72.866Zm30.83 16.292V172.754L0 115.906l65.32-38.532l91.541 55.952v104.53L93.846 276.78Zm90.932-12.793V119.229L57.516 41.861L128.24 0L256 74.049V221.17l-71.221 42.816Z"/>`),
		g.Group(children),
	)
}