package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.78 5.5H5.5v8.28m37 0V5.5h-8.28m0 37h8.28v-8.28m-31.95-7.81v11h11v-11h-11Zm3.22 3.22h4.6v4.6h-4.6Zm-3.22-19.08v11h11v-11h-11Zm3.22 3.22h4.6v4.6h-4.6Zm12.64-3.22v11h11v-11h-11Zm3.22 3.22h4.6v4.6h-4.6Zm2.3 23.68a5.79 5.79 0 0 0 3.47-2.1a8.37 8.37 0 0 0 2.05-5.72v-3.22h-11v3.22h0a8.39 8.39 0 0 0 2.05 5.72a5.79 5.79 0 0 0 3.43 2.1Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.93 34.23a5.82 5.82 0 0 1-1.08-.86a5.4 5.4 0 0 1-1.22-3.74h4.6A5.4 5.4 0 0 1 33 33.37a5.82 5.82 0 0 1-1.08.86ZM5.5 34.22v8.28h8.28"/>`),
		g.Group(children),
	)
}