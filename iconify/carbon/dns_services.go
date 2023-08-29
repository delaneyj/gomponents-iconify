package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnsServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 20h10v2H20zm0 4h6v2h-6z"/><path fill="currentColor" d="M30 17v-1a13.987 13.987 0 1 0-10.77 13.625l-.46-1.946A12.042 12.042 0 0 1 16 28c-.19 0-.375-.019-.563-.027A20.304 20.304 0 0 1 12.026 17Zm-2.042-2h-5.983a24.284 24.284 0 0 0-2.774-10.559A12.023 12.023 0 0 1 27.96 15ZM16.564 4.027A20.304 20.304 0 0 1 19.974 15h-7.948a20.304 20.304 0 0 1 3.411-10.973C15.625 4.02 15.81 4 16 4s.375.019.563.027Zm-3.764.414A24.284 24.284 0 0 0 10.025 15H4.04A12.023 12.023 0 0 1 12.8 4.441Zm0 23.118A12.023 12.023 0 0 1 4.042 17h5.983a24.284 24.284 0 0 0 2.774 10.559Z"/>`),
		g.Group(children),
	)
}