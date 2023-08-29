package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SrfNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M21.083 22.845V14.04h2.883c1.63 0 2.951 1.324 2.951 2.958s-1.321 2.957-2.951 2.957l2.883 2.89m-14.355-.965c.54.704 1.217.965 2.159.965h1.304a2.197 2.197 0 0 0 2.197-2.196v-.01a2.197 2.197 0 0 0-2.197-2.197h-1.438a2.2 2.2 0 0 1-2.2-2.2h0c0-1.216.987-2.203 2.204-2.203h1.297c.942 0 1.62.262 2.16.965m12.429 3.438h2.863m-2.863 4.403v-8.806h4.404"/><path d="M8.5 9.908h31v17.069h-31z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.469 32.792l-1.65 5.3l-1.65-5.3l-1.65 5.3l-1.65-5.3m8.324 4.853c.366.307.76.447 1.645.447h.449c.73 0 1.322-.593 1.322-1.325h0c0-.732-.592-1.325-1.322-1.325h-.897c-.73 0-1.323-.593-1.323-1.325h0c0-.732.592-1.325 1.323-1.325h.448c.886 0 1.28.14 1.645.447m-19.092 4.853v-3.3a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2m0 3.3v-5.3m9.624 4.291a2 2 0 0 1-1.738 1.01h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v.65h-4"/>`),
		g.Group(children),
	)
}