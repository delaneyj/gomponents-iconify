package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Funktrainer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.093 31.892a15.815 15.815 0 0 1 0-27.392"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.674 29.153a12.652 12.652 0 0 1 0-21.914"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.256 26.414a9.49 9.49 0 0 1 0-16.436"/><circle cx="24" cy="18.196" r="2.109" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5V20.305M31.907 4.5a15.815 15.815 0 0 1 0 27.392"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.326 7.24a12.652 12.652 0 0 1 0 21.913"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.744 9.978a9.49 9.49 0 0 1 0 16.436"/>`),
		g.Group(children),
	)
}