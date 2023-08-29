package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDatabaseServerTwoServerNetworkInternet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<rect width="13" height="5" x=".5" y=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="currentColor" d="M3.25 2.25A.75.75 0 1 0 4 3a.76.76 0 0 0-.75-.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 3h2M1.5 5.5a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1"/><path fill="currentColor" d="M3.25 7.25A.75.75 0 1 0 4 8a.76.76 0 0 0-.75-.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 8h2m-1 2.5v3m-5 0h10"/>`),
		g.Group(children),
	)
}