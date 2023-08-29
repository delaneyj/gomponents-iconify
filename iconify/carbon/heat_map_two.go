package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatMapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="21" cy="20" r="2" fill="currentColor"/><circle cx="14" cy="12" r="2" fill="currentColor"/><circle cx="29" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M26.5 30a3.5 3.5 0 1 1 3.5-3.5a3.504 3.504 0 0 1-3.5 3.5zm0-5a1.5 1.5 0 1 0 1.5 1.5a1.502 1.502 0 0 0-1.5-1.5zM14 30a3.958 3.958 0 0 1-2.126-.621a6.998 6.998 0 1 1 4.11-6.838A3.992 3.992 0 0 1 14 30zm-1.884-3.028l.539.495a1.992 1.992 0 1 0 2.004-3.343l-.691-.243l.03-.847a5.008 5.008 0 1 0-2.517 4.302zM24 16a6.007 6.007 0 0 1-6-6a5.325 5.325 0 0 1 .027-.533A3.956 3.956 0 0 1 16 6a4.005 4.005 0 0 1 4-4a3.956 3.956 0 0 1 3.467 2.027C23.648 4.01 23.825 4 24 4a6 6 0 0 1 0 12zM20 4a2.002 2.002 0 0 0-2 2a1.98 1.98 0 0 0 1.43 1.902l.902.27l-.215.917A3.994 3.994 0 1 0 24 6a4.006 4.006 0 0 0-.912.116l-.916.214l-.27-.9A1.98 1.98 0 0 0 20 4zM6.5 11A4.5 4.5 0 1 1 11 6.5A4.505 4.505 0 0 1 6.5 11zm0-7A2.5 2.5 0 1 0 9 6.5A2.503 2.503 0 0 0 6.5 4z"/>`),
		g.Group(children),
	)
}