package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthDataSecurityNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsHealthDataSecurityNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM20 4a3 3 0 0 0-3 3h-3a3 3 0 0 0-3 3v28a3 3 0 0 0 3 3h12v-2H14a1 1 0 0 1-1-1V10a1 1 0 0 1 1-1h3a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3h3a1 1 0 0 1 1 1v15h2V10a3 3 0 0 0-3-3h-3a3 3 0 0 0-3-3h-8Zm-1 3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1V7Zm4 12v-3h2v3h3v2h-3v3h-2v-3h-3v-2h3Zm4 9a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v6.789a7 7 0 0 1-3.117 5.824l-3.328 2.22a1 1 0 0 1-1.11 0l-3.328-2.22A7 7 0 0 1 27 34.79V28Zm6.707 8.707l4-4l-1.414-1.414L33 34.586l-1.293-1.293l-1.414 1.414l2 2a1 1 0 0 0 1.414 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHealthDataSecurityNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}