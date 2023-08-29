package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PdfScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 10V6.5c0-1.1-.9-2-2-2H34M39.5 38v3.5c0 1.1-.9 2-2 2H34"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M35.923 17.462H29c-.85 0-1.538-.69-1.538-1.539V9H13.615c-.85 0-1.538.689-1.538 1.538v26.924c0 .85.689 1.538 1.538 1.538h20.77c.85 0 1.538-.689 1.538-1.538v-20ZM27.462 9l8.461 8.462m-7.354 6.949h3.455m-3.455 3.449h2.254m-2.254-3.449v6.897"/><path d="M15.976 31.294v-6.896h2.321a2.314 2.314 0 1 1 .002 4.629h-2.323m6.297 2.281v-6.923h1.174a3.461 3.461 0 0 1 3.461 3.461h0a3.461 3.461 0 0 1-3.462 3.462h-1.173Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 10V6.5c0-1.1.9-2 2-2H14m0 39h-3.5c-1.1 0-2-.9-2-2V38"/>`),
		g.Group(children),
	)
}