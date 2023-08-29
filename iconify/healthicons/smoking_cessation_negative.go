package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmokingCessationNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsSmokingCessationNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM28.727 37H7a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2h17.62a2 2 0 0 1 1.474.649l3.37 3.675c.588.642.133 1.676-.737 1.676ZM10 8.586l.707.707l28 28l.707.707L38 39.414l-.707-.707l-28-28L8.586 10L10 8.586ZM40 32a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2Zm-6-15a3 3 0 0 0-3 3v1.818a3 3 0 0 0 3 3h3a3 3 0 0 1 3 3V28a1 1 0 1 1-2 0v-.182a1 1 0 0 0-1-1h-3a5 5 0 0 1-5-5V20a5 5 0 0 1 5-5a1 1 0 1 1 0 2Zm3 3a1 1 0 0 0 1 1a4 4 0 0 1 4 4v3a1 1 0 1 0 2 0v-3a5.994 5.994 0 0 0-2.644-4.974A4.4 4.4 0 0 0 43 16.59V15a5 5 0 0 0-5-5a1 1 0 1 0 0 2a3 3 0 0 1 3 3v1.59A2.41 2.41 0 0 1 38.59 19H38a1 1 0 0 0-1 1Zm7 12a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsSmokingCessationNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}