package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassengerPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.5 9A3.5 3.5 0 1 1 14 5.5A3.504 3.504 0 0 1 10.5 9zm0-5A1.5 1.5 0 1 0 12 5.5A1.502 1.502 0 0 0 10.5 4zm11.974 27.313L19.34 24h-7.101a4.007 4.007 0 0 1-3.867-2.97l-1.634-6.127a3.899 3.899 0 0 1 7.535-2.009L15.1 16H21v2h-7.436l-1.223-4.59a1.9 1.9 0 0 0-3.67.978l1.633 6.126A2.005 2.005 0 0 0 12.238 22h8.42l3.655 8.525zM30 6h-4V2h-2v4h-4v2h4v4h2V8h4V6z"/><path fill="currentColor" d="M18 28H7.768a2.002 2.002 0 0 1-1.933-1.485L2.033 12.258l1.933-.516L7.768 26H18Z"/>`),
		g.Group(children),
	)
}