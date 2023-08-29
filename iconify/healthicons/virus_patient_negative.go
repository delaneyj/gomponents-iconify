package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirusPatientNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsVirusPatientNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM17 18c3.867 0 7-3.133 7-7s-3.133-7-7-7s-7 3.133-7 7s3.133 7 7 7Zm0 1c-4.339 0-13 2.144-13 6.4V31h16.153a13.013 13.013 0 0 1 6.826-9.525C24.06 19.828 19.696 19 17 19Zm14 3a1 1 0 1 0 0 2h1v2.07a6.965 6.965 0 0 0-3.192 1.324L27.414 26l.293-.293a1 1 0 0 0-1.414-1.414l-2 2a1 1 0 0 0 1.414 1.414l.293-.293l1.394 1.394A6.965 6.965 0 0 0 26.07 32H24v-1a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-1h2.07a6.965 6.965 0 0 0 1.324 3.192L26 38.586l-.293-.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414-1.414L27.414 40l1.394-1.394A6.965 6.965 0 0 0 32 39.93V42h-1a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-1v-2.07a6.975 6.975 0 0 0 3.738-1.777L39.586 40l-.293.293a1 1 0 0 0 1.414 1.414l2-2a1 1 0 0 0-1.414-1.414l-.293.293l-1.992-1.992c.467-.78.787-1.657.921-2.594H42v1a1 1 0 1 0 2 0v-4a1 1 0 1 0-2 0v1h-2.07a6.965 6.965 0 0 0-1.324-3.192L40 27.414l.293.293a1 1 0 0 0 1.414-1.414l-2-2a1 1 0 0 0-1.414 1.414l.293.293l-1.394 1.394A6.965 6.965 0 0 0 34 26.07V24h1a1 1 0 1 0 0-2h-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsVirusPatientNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}