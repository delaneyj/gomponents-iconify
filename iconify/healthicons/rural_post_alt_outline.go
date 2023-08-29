package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuralPostAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M25 9.548a1 1 0 1 0-2 0v2.5h-2.5a1 1 0 1 0 0 2H23v2.5a1 1 0 0 0 2 0v-2.5h2.5a1 1 0 0 0 0-2H25v-2.5Z"/><path fill-rule="evenodd" d="M25 37v-8h11v8H25Zm2-2v-4h7v4h-7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 10.048a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v5h8.388c.706 0 1.318.487 1.477 1.175l1.11 4.822A4.899 4.899 0 0 1 40 26.17v15.877H8V26.171a4.899 4.899 0 0 1-1.976-5.126l1.111-4.822a1.515 1.515 0 0 1 1.477-1.175H17v-5Zm10-2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h6Zm0 12a4.002 4.002 0 0 0 3.874-3h1.267v4.469a3.531 3.531 0 0 1-7.063 0v-1.47H27Zm-9.874-3c.444 1.725 2.01 3 3.874 3h1.922v1.469a3.531 3.531 0 0 1-7.063 0v-4.47h1.267Zm20.073 10c.274 0 .541-.022.801-.065v13.065H22V29h-9v11.048h-3V26.983a4.898 4.898 0 0 0 4.93-2.195a5.524 5.524 0 0 0 4.46 2.26A5.526 5.526 0 0 0 24 24.575a5.526 5.526 0 0 0 4.61 2.473c1.83 0 3.453-.89 4.46-2.26a4.898 4.898 0 0 0 4.129 2.26Zm-28.202-10l-1.024 4.446a2.902 2.902 0 1 0 5.73.652v-5.098H8.997Zm31.03 4.446l-1.025-4.446h-4.705v5.098a2.902 2.902 0 1 0 5.73-.652ZM20 31h-5v9h5v-9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}