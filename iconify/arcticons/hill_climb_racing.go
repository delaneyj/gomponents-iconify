package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HillClimbRacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 17.331c-9.01 11.656-18.88 22.202-37 22.087"/><circle cx="33.817" cy="20.284" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="20.532" cy="32.724" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.267 29.524l7.73-8.407c-3.688-3.352.712-8.209 4.417-5.34l-.112-.97l.338-.518l-3.47-3.584"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.037 11.642c-2.353-.786-4.306 1.59-6.21 3.075l-4.01-.18m4.011.18L12.756 28.915l.068 1.488l3.38 3.358l-.63.63"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.204 33.76l.97.52c-3.456-3.868 1.848-8.412 5.157-4.825M9.015 22.943c2.827.123 4.976 1.205 7.096 2.329m2.384-2.59l-1.615-.01m3.852-2.42c-.617-.764-2.2-.947-1.29-2.626m-.15 4.19c-1.002-.187-1.311-2.767-2.463-.971"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.442 17.625c-4.782-4.706-9.792 1.864-2.613 3.22m.391-4.667l.565-.752m2.611 7.269l-.834.924m-3.808-8.316l-.97 6.423"/>`),
		g.Group(children),
	)
}