package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.89 26A1.94 1.94 0 0 0 26 27.89v12.66a1.94 1.94 0 0 0 1.94 2h12.61a2 2 0 0 0 2-2V27.89a1.94 1.94 0 0 0-2-1.94ZM7.45 5.5a2 2 0 0 0-1.95 2v12.61a1.94 1.94 0 0 0 2 1.94h12.61a1.94 1.94 0 0 0 1.94-1.94V7.45a1.94 1.94 0 0 0-1.94-1.95Z"/><rect width="16.55" height="16.55" x="25.95" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.92"/><rect width="16.55" height="16.55" x="5.5" y="25.95" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.95"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.78 8.42v10.71m-5.36-5.35h10.71m9.74 17.52h10.71m-10.71 5.84h10.71"/>`),
		g.Group(children),
	)
}