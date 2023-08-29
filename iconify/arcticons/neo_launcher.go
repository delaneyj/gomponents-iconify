package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeoLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.144 30.894A9.75 9.75 0 1 1 14.25 14.25c5.385 0 7.85 5.287 9.75 9.75c2.11 4.954 4.365 9.75 9.75 9.75a9.75 9.75 0 1 0-6.894-16.644M7.75 24h13m-11-4h9m-9 8h9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.153 27.918v.48a.603.603 0 0 1-.603.602h-1.941a.603.603 0 0 1-.603-.603v-.63a.598.598 0 0 1 .385-.56a4.267 4.267 0 0 0 2.762-3.95a4.405 4.405 0 0 0-8.806 0a4.267 4.267 0 0 0 2.762 3.95a.598.598 0 0 1 .385.56v.63a.603.603 0 0 1-.603.603H29.95a.603.603 0 0 1-.603-.603v-.48"/>`),
		g.Group(children),
	)
}