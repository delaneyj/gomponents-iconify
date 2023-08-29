package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M31.352 17.666L25 13.463V11h3V9h-3V6h-2v3h-3v2h3v2.463l-6.352 4.203a1 1 0 0 0-.448.834v4.377l-9.56 3.69A1 1 0 0 0 6 27.5V42h2V28.186l9.56-3.69a1 1 0 0 0 .64-.933v-4.526L24 15.2l5.8 3.838v4.526a1 1 0 0 0 .734.963L40 27.136V42h2V26.375a1 1 0 0 0-.734-.964l-9.466-2.61V18.5a1 1 0 0 0-.448-.834Z"/><path d="M17 30a1 1 0 1 0-2 0v12h2V30Zm15-1a1 1 0 0 1 1 1v12h-2V30a1 1 0 0 1 1-1Zm-8 1a3 3 0 0 0-3 3v9h6v-9a3 3 0 0 0-3-3Z"/></g>`),
		g.Group(children),
	)
}