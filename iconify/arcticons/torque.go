package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torque(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.413 30.857v-7.529h1.694a3.294 3.294 0 0 1 3.293 3.294v.941a3.294 3.294 0 0 1-3.293 3.294ZM14.6 28.363a2.494 2.494 0 0 0 4.987 0v-2.54a2.494 2.494 0 0 0-4.987 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M24.696 27.093a1.882 1.882 0 0 1 0 3.764h-3.105v-7.529h3.105a1.882 1.882 0 0 1 0 3.764Zm0 0h-3.105"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.142 14.703a1.7 1.7 0 0 1 0-3.39h12a1.7 1.7 0 1 1 0 3.39ZM7.893 31.573a1.7 1.7 0 0 1-3.39 0h0v-10.17a1.7 1.7 0 1 1 3.39 0Zm0-5.09h2.109m14.14-9.021v-2.759m12.631 10.36h2.3v-3.26h2.11l2.32 7.2l-2.33 7.2h-2.1v-3.14h-2.3l.02 2.05a1.58 1.58 0 0 1-1.58 1.58h-13.56a1.17 1.17 0 0 1-.66-.2l-4.99-3.47h-4.21a1.79 1.79 0 0 1-1.79-1.79v-.041h0v-9.4a1.8 1.8 0 0 1 1.79-1.77h1.4l2.66-2.26a1.13 1.13 0 0 1 .73-.27h15.26a1.59 1.59 0 0 1 1.58 1.53v2.15h1.79a1.59 1.59 0 0 1 1.58 1.59h0Z"/>`),
		g.Group(children),
	)
}