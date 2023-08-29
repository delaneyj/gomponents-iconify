package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebInspectorFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.853 13.657a21.479 21.479 0 1 1-9.87-9.197"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.856 27.826s3.04-8.366-2.621-12.74c-2.886-2.23-3.581-5.727-2.128-8.244a6.762 6.762 0 0 1 9.237-2.475c3.234 1.867 3.639 5.946 2.475 9.237c-4.643 9.779-11.439 25.904-17.902 24.845c-3.79-.62-.593-14.158-.593-14.158c-4.698 13.881-12.402 17.496-14.613 15.774s-1.605-7.031 1.139-12.22a30.671 30.671 0 0 0 3.018-11.393"/>`),
		g.Group(children),
	)
}