package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapleCalculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.137 30.794s1.476 1.677 1.476 3.744s-1.476 3.743-1.476 3.743m-3.877-7.487s-1.475 1.677-1.475 3.744s1.475 3.743 1.475 3.743m3.365-5.709l-2.951 3.909m2.951 0l-2.951-3.909m24.81-19.82l-3.365 4.459m3.365 0l-3.365-4.459"/><circle cx="24" cy="24" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5V17m7 7h11.5M24 31v11.5M17 24H5.5m1.588-9.576h2.05l1.916 4.575l3.223-9.668h8.114M12.252 30.866c-2.355.011-2.176 1.242-2.614 3.346c-.439 2.102-.085 3.924-2.66 4.069m1.262-4.248h2.66m27.887-2.098h-6.92m4.598 0l.045 3.81c-.024 2.515 1.763 2.263 2.322 1.08m-4.823-4.89c-.474 1.913.644 5.79-1.488 5.657"/>`),
		g.Group(children),
	)
}