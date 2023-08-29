package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orgro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.392 12.332a24.883 24.883 0 0 0-2.577.13l.05-.012s-2.59-3.009-5.427-2.728a6.81 6.81 0 0 1 1.45 3.626l.058-.013a10.86 10.86 0 0 0-.771.33l.018-.015L8.1 4.5l7.068 12.6l.011-.01c-.342.46-.665.937-.982 1.426c-2.154 3.329-6.412 10.793-6.412 10.793l3.154 3.048h4.783c0-1.892.736-4.52 4.993-5.571s5.256-2.733 5.571-4.783c0 0 .579.683.579 2.996S24 28.94 24 28.94a12.238 12.238 0 0 1 3.863.683a7.99 7.99 0 0 1 2.68 1.682a19.233 19.233 0 0 0 1.157-5.782c0-2.733-.946-7.79-6.623-9.493c0 0 10.88-1.674 10.88 10.965a19.233 19.233 0 0 1-1.997 8.094a25.002 25.002 0 0 1 2.628 8.41s3.627-7.674 3.627-15.085S37.009 12.33 26.391 12.33Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.95 23.036c1.436-1.261 3.574-3.574 4.03-5.01l-.35 2.277l.35.631a4.471 4.471 0 0 1-4.03 2.102Z"/>`),
		g.Group(children),
	)
}