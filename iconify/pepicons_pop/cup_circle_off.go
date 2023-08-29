package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M18.5 7h-12a1 1 0 0 0-1 1c0 4.918 3.061 9 7 9s7-4.082 7-9a1 1 0 0 0-1-1Zm-6 8c-2.455 0-4.596-2.57-4.949-6h9.898c-.353 3.43-2.494 6-4.949 6Z" clip-rule="evenodd"/><path d="M7 17.5h11a1 1 0 1 1 0 2H7a1 1 0 1 1 0-2Zm10.024-3.69l.552-1.923c.257.074.539.113.831.113c1.107 0 1.893-.543 1.893-1c0-.457-.786-1-1.893-1V8c2.088 0 3.893 1.248 3.893 3s-1.805 3-3.893 3c-.477 0-.945-.065-1.383-.19Z"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}