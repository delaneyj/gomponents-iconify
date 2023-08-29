package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotfindskittenAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.83 21.358a1.386 1.386 0 0 0-1.076 2.26l2.222 2.609l2.2-2.583l.01-.012l.012-.014a1.387 1.387 0 1 0-2.222-1.653a1.383 1.383 0 0 0-1.145-.607h-.001Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.001" d="M10.108 17.461c3.095 0 5.604 2.475 5.604 5.528v.79H4.505v-.79c0-3.053 2.509-5.528 5.604-5.528h0Zm-4.839-.826l1.592 1.849m8.086-1.849l-1.592 1.849"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12.657" cy="21.23" r="1.02"/><circle cx="7.559" cy="21.23" r="1.02"/><circle cx="39.167" cy="24.795" r="1.02"/><circle cx="34.069" cy="24.795" r="1.02"/><path d="m31.01 16.638l4.078 5.098h3.06l4.077-5.098v14.274H31.01V16.639Z"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M29.735 28.873h2.549m-2.549-1.529h2.549m8.667 1.529H43.5m-2.549-1.529H43.5m-6.882 1.019v-1.332m0 1.332l-.805.998m-.361-1.73l1.166.732m0 0l.804.998m.361-1.73l-1.165.732"/><path stroke-width=".993" d="M4.5 23.774h11.216v5.608H4.5z"/></g><circle cx="7.608" cy="30.615" r=".75" fill="currentColor"/><circle cx="12.608" cy="30.615" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}