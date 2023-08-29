package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Krita(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.587 7.196a21.5 21.5 0 1 1-3.226 3.188"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.08 26.692s-2.535.834-2.59 2.609m-4.55-3.644s-.848-4.276 3.322-3.433"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.411 37.26c9.106-6.543.671-10.563.671-10.563s-3.64-3.983-3.823-4.473c-.167-.448.218-2.485-5.85-7.614S3.766 2.544 3.766 2.544a1.138 1.138 0 0 0-1.034.21a1.138 1.138 0 0 0-.176 1.04s7.267 10.35 12.585 16.253s7.343 5.454 7.796 5.607c.495.167 4.552 3.647 4.552 3.647a4.66 4.66 0 0 0 1.444 3.675s2.562 1.998.478 4.285Z"/>`),
		g.Group(children),
	)
}