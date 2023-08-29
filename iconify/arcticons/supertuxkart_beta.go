package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SupertuxkartBeta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.43 43.327A21.413 21.413 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24a21.4 21.4 0 0 1-2.174 9.431"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.947 19.96a4.536 4.536 0 0 1 2.1 2.092l10.702-2.81a15.529 15.529 0 0 0-9.99-9.99L25.948 19.96Zm-5.987 2.093a4.536 4.536 0 0 1 2.092-2.1L19.241 9.252a15.529 15.529 0 0 0-9.99 9.989l10.709 2.812Zm17.56 9.516a15.4 15.4 0 0 0 1.228-2.81l-10.707-2.812a4.536 4.536 0 0 1-2.093 2.1l2.811 10.702a15.52 15.52 0 0 0 2.81-1.228m-9.517-9.474a4.536 4.536 0 0 1-2.092-2.1L9.251 28.759a15.529 15.529 0 0 0 9.989 9.99l2.81-10.703Z"/><circle cx="24" cy="24" r="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.536 38.5a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4h0Zm0 0h-3.3"/>`),
		g.Group(children),
	)
}