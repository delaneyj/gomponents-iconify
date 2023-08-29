package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchPlayBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.25 2.75A.75.75 0 0 1 10 2h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75Zm11.75 11a9 9 0 1 1-18 0a9 9 0 0 1 18 0Zm-7.974-2.725a14.502 14.502 0 0 0-.784-.508c-1.073-.652-1.609-.978-2.09-.617c-.48.36-.524 1.116-.612 2.628c-.024.427-.04.846-.04 1.222s.016.795.04 1.222c.088 1.512.132 2.267.612 2.628c.481.361 1.018.035 2.09-.617c.278-.169.547-.341.784-.508c.27-.19.565-.418.862-.66c1.075-.877 1.612-1.315 1.612-2.065s-.537-1.188-1.612-2.065c-.297-.242-.591-.47-.862-.66Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}