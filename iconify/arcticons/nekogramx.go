package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nekogramx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.08 42.5c-.63-1.88-1.32-3.54-.86-5.74c-1.72-2.24-3.31-9.26-1.89-14.82c3.16.8 6.25 1.78 8.83 4.09a16 16 0 0 1 7.23-4.68a24.51 24.51 0 0 1-1.34-10.5c3.86-.28 7.09 1.76 9.37 4"/><ellipse cx="27.66" cy="34.93" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.71" ry="1.88" transform="rotate(-34.88 27.654 34.935)"/><ellipse cx="38.36" cy="26.38" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.57" ry="1.73" transform="rotate(-47.47 38.353 26.376)"/>`),
		g.Group(children),
	)
}