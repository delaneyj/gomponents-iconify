package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nests(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTNests0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4.999 32.314L20.034 5.943C20.958 3.98 22.627 3 25.04 3c3.62 0 5.977 4.987 5.977 4.987L35 8.58c-4.01.065-6.33.872-6.957 2.42c-.94 2.322 2.456 4.731 2.975 8.004c.52 3.273-1.55 8.801-6.529 11.563c-3.319 1.841-7.462 2.32-12.43 1.433l-6.1 11"/><path fill="#555" fill-rule="evenodd" d="M14.942 26.978c2.58-4.65 4.275-7.635 5.087-8.956c1.219-1.982 6.245-1.385 4.44 3.563c-1.204 3.298-4.38 5.096-9.527 5.393Zm18.026 6.006c1.142-2.737 2.81-4.395 5.003-4.975c2.193-.58 4.193-.227 6 1.06c-1.125 2.595-2.794 4.26-5.005 4.992c-2.211.733-4.21.373-5.998-1.077Z" clip-rule="evenodd"/><path fill="#555" fill-rule="evenodd" d="M26.985 35.114c.757 2.722 2.762 4.543 6.014 5.461c3.253.92 5.93.394 8.033-1.575c-2.6-3.03-4.944-4.694-7.03-4.99c-2.087-.296-4.426.072-7.017 1.104Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTNests0)"/>`),
		g.Group(children),
	)
}