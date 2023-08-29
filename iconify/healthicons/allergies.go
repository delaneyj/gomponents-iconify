package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Allergies(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M31 25.731V9a2 2 0 1 0-4 0v14h-2V6a2 2 0 1 0-4 0v17h-2V9a2 2 0 1 0-4 0v16.111h-2V15a2 2 0 1 0-4 0v17c0 6.627 5.373 12 12 12h1.422a12 12 0 0 0 10.553-6.287l6.938-12.816a2.552 2.552 0 0 0-4.36-2.642L32.83 26.29c-.552.817-1.829.427-1.829-.56Zm-12.977 4.252a1 1 0 1 0-.002-2a1 1 0 0 0 .002 2Zm5 .996a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-1.997 3.002a1 1 0 1 1-2 .002a1 1 0 0 1 2-.002Zm4-.003a1 1 0 1 0-.001-2a1 1 0 0 0 .001 2Zm3.003 2.997a1 1 0 1 1-2 .002a1 1 0 0 1 2-.002Zm.996-4a1 1 0 1 0-.001-2a1 1 0 0 0 .001 2Zm-5.996 5.004a1 1 0 1 1-2 .002a1 1 0 0 1 2-.002Zm-6 .005a1 1 0 1 0-.001-2a1 1 0 0 0 .001 2Zm-1.004-4.999a1 1 0 1 1-2 .002a1 1 0 0 1 2-.002Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}