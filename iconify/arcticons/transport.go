package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.54 39.08h7.74v4.42H8.54zm23.18 0h7.74v4.42h-7.74zM7.29 11h1.2v7.94a2.44 2.44 0 0 1-2.4-2.47v-4.25A1.22 1.22 0 0 1 7.29 11Zm33.42 0h-1.2v7.94a2.44 2.44 0 0 0 2.4-2.47v-4.25a1.22 1.22 0 0 0-1.2-1.22ZM24 4.5c-5.51 0-15.46.86-15.46 3.1v31.48h30.92V7.6c0-2.24-10.02-3.1-15.46-3.1Zm0 4.23q8.19 0 12.2 1.23a1 1 0 0 1 .67.92v17a1 1 0 0 1-1 1H12.14a1 1 0 0 1-1-1v-17a1 1 0 0 1 .67-.92q4-1.23 12.19-1.23ZM13.59 31.5a2.5 2.5 0 0 1 0 5h0a2.47 2.47 0 0 1-2.43-2.5h0a2.47 2.47 0 0 1 2.43-2.5Zm20.87 0a2.5 2.5 0 0 1 0 5h0A2.47 2.47 0 0 1 32 34h0a2.47 2.47 0 0 1 2.46-2.5Z"/>`),
		g.Group(children),
	)
}