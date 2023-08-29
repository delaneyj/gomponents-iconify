package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirusResearchNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsVirusResearchNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM21 6a1 1 0 1 0 0 2h1v3.041a11.948 11.948 0 0 0-6.75 2.796l-.028-.03l-2.121-2.12l.707-.708a1 1 0 0 0-1.415-1.414l-2.828 2.828a1 1 0 1 0 1.414 1.415l.707-.707l2.122 2.12l.03.03A11.947 11.947 0 0 0 11.04 22H8v-1a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-1h3.041a11.947 11.947 0 0 0 2.796 6.75l-.03.028l-2.12 2.122l-.708-.707a1 1 0 0 0-1.414 1.414l2.828 2.828a1 1 0 1 0 1.415-1.414l-.707-.707l2.12-2.121l.03-.03a11.948 11.948 0 0 0 7.203 2.825A9.996 9.996 0 0 1 22 32c0-5.523 4.477-10 10-10c1.04 0 2.044.159 2.988.454a11.948 11.948 0 0 0-2.825-7.204l.03-.028l2.12-2.121l.708.707a1 1 0 0 0 1.414-1.415l-2.828-2.828a1 1 0 1 0-1.415 1.414l.708.707l-2.122 2.122l-.028.03A11.948 11.948 0 0 0 24 11.04V8h1a1 1 0 1 0 0-2h-4Zm1 14a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm10 6a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm-8 6a8 8 0 1 1 14.32 4.906l4.387 4.387a1 1 0 0 1-1.414 1.414l-4.387-4.387A8 8 0 0 1 24 32Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsVirusResearchNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}