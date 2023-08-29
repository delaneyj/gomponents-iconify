package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m0 600l206.909 206.909l-.623-144.765h331.567v332.153l-142.933.623L600 1200l206.909-206.909l-144.765.622V662.146H994.3l.622 142.933L1200 600L993.091 393.091l.623 144.763H662.146V205.701l142.933-.623L600 0L393.091 206.909l144.765-.623v331.567H205.701l-.623-142.933L0 600z"/>`),
		g.Group(children),
	)
}