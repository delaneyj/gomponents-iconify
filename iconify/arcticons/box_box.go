package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.975 22.487c.501-.25.501-.74.501-.987v-.656"/><circle cx="35.225" cy="20.75" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.321 13l5.3 8m0-8l-5.3 8m-4.996 0a2.65 2.65 0 0 1-2.65-2.65v-2.7a2.65 2.65 0 1 1 5.3 0v2.7a2.65 2.65 0 0 1-2.65 2.65Zm-7-4a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4Zm0 0h-3.3m15.296 9.5l5.3 8m0-8l-5.3 8m-4.996 0a2.65 2.65 0 0 1-2.65-2.65v-2.7a2.65 2.65 0 1 1 5.3 0v2.7a2.65 2.65 0 0 1-2.65 2.65Zm-7-4a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4Zm0 0h-3.3"/><circle cx="35.225" cy="34.25" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.225 26.25v5.95m-21.55 6.4h15.35M13.675 9.4h20.3"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}