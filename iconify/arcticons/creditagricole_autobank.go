package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditagricoleAutobank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.325 29.917h30.08M21.132 10.405c-7.16.885-12.178 3.509-12.178 9.183c0 3.462 3.563 5.161 6.844 5.161c6.804 0 17.15-6.567 23.978-10.416"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.187 20.559l5.446-10.154l7.832 15.047m5.311-16.193a68.023 68.023 0 0 1-7.192 4.047M7.776 37.068h-2.23m-.558 1.673l1.673-4.957l1.673 4.957m25.073-1.673h-2.23m-.557 1.673l1.672-4.957l1.673 4.957m-19.339-4.957h3.223m-1.611 4.957v-4.957m-6.405.001v3.284c0 .929.744 1.673 1.611 1.673s1.611-.744 1.611-1.673v-3.284m27.355-.001v4.957m0-1.735l2.602-3.222m0 4.957l-1.982-2.478m-5.651 2.478v-4.957l3.284 4.957v-4.957m-18.047 4.958a1.6 1.6 0 0 1-1.612-1.611v-1.673c0-.929.744-1.673 1.612-1.673c.929 0 1.672.744 1.672 1.673v1.611c0 .929-.743 1.673-1.672 1.673Zm7.545-2.478c.682 0 1.24.558 1.24 1.239c0 .682-.558 1.239-1.24 1.239h-2.044v-4.957h2.044c.682 0 1.24.558 1.24 1.239s-.558 1.24-1.24 1.24Zm0-.001h-2.044"/>`),
		g.Group(children),
	)
}