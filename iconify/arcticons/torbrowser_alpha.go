package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TorbrowserAlpha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.326 33.432a21.486 21.486 0 1 0-9.894 9.894"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 19.01a4.99 4.99 0 0 1 0 9.98"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 13.725a10.275 10.275 0 0 1 0 20.55m7.505 3.997A16.052 16.052 0 0 1 24 40.122"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.878a16.127 16.127 0 0 1 14.271 23.628"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.9 42.476l2.599-7.976m2.601 8l-2.601-8m1.731 5.324h-3.466"/>`),
		g.Group(children),
	)
}