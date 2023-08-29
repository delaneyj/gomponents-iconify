package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DischargeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsDischargeNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM23 10v3h-3v2h3v3h2v-3h3v-2h-3v-3h-2ZM4.064 21.351a1 1 0 1 1 1.873-.702l.275.735a1 1 0 0 1 1.737.3L8.72 24h9.281v2h-2v2h1a1 1 0 0 1 .04 2l1.714 6h10.491l1.715-6a1 1 0 0 1 .04-2h1v-2h-2v-2h9.279l.772-2.316a1 1 0 0 1 1.736-.3l.276-.735a1 1 0 0 1 1.873.702l-2.891 7.709L42.78 36H43a1 1 0 1 1 0 2H5a1 1 0 1 1 0-2h.22l1.735-6.94l-2.891-7.709ZM39.307 28H34v-2h6.057l-.75 2ZM14 28v-2H7.943l.75 2H14Zm26.72 8h-9.394l1.429-5h6.714l1.25 5Zm-24.046 0l-1.428-5H8.531l-1.25 5h9.393Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDischargeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}