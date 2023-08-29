package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirusAltNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsVirusAltNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM20.4 5.9a1 1 0 1 0 0 2H23v2.636a13.442 13.442 0 0 0-7.813 3.237l-1.923-1.923l1.543-1.543a1 1 0 1 0-1.414-1.414l-4.5 4.5a1 1 0 1 0 1.414 1.414l1.543-1.543l1.923 1.923A13.443 13.443 0 0 0 10.537 23H7.9v-2.6a1 1 0 1 0-2 0v7.2a1 1 0 1 0 2 0V25h2.636a13.443 13.443 0 0 0 3.237 7.813l-1.923 1.923l-1.543-1.543a1 1 0 0 0-1.414 1.414l4.5 4.5a1 1 0 0 0 1.414-1.414l-1.543-1.543l1.923-1.923A13.442 13.442 0 0 0 23 37.463V40.1h-2.6a1 1 0 1 0 0 2h7.2a1 1 0 1 0 0-2H25v-2.637a13.442 13.442 0 0 0 7.813-3.236l1.923 1.923l-1.543 1.543a1 1 0 0 0 1.414 1.414l4.5-4.5a1 1 0 1 0-1.414-1.414l-1.543 1.543l-1.923-1.923A13.442 13.442 0 0 0 37.464 25H40.1v2.6a1 1 0 1 0 2 0v-7.2a1 1 0 1 0-2 0V23h-2.636a13.442 13.442 0 0 0-3.237-7.813l1.923-1.923l1.543 1.543a1 1 0 0 0 1.414-1.414l-4.5-4.5a1 1 0 1 0-1.414 1.414l1.543 1.543l-1.923 1.923A13.442 13.442 0 0 0 25 10.536V7.9h2.6a1 1 0 1 0 0-2h-7.2Zm11.7 17.65a1.35 1.35 0 1 1-2.7 0a1.35 1.35 0 0 1 2.7 0Zm-7.65 8.55a1.35 1.35 0 1 0 0-2.7a1.35 1.35 0 0 0 0 2.7ZM23.1 20.4a2.7 2.7 0 1 1-5.4 0a2.7 2.7 0 0 1 5.4 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsVirusAltNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}