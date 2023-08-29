package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudProjects(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 18H20v6.468a5.022 5.022 0 0 0 2.861 4.52L25 30l2.139-1.013A5.022 5.022 0 0 0 30 24.467V18zm-5 9.786l-1.283-.607A3.012 3.012 0 0 1 22 24.468V20h6v4.468a3.012 3.012 0 0 1-1.717 2.71L25 27.787zM17 18H5a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h12v-2H5v-5h12v-2zM27 4h-5a2 2 0 0 0-2 2v9h2V6h5v9h2V6a2 2 0 0 0-2-2zM15 4H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zM5 13V6h10v7H5z"/>`),
		g.Group(children),
	)
}