package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeMiniOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19q-2.9 0-4.95-2.05T2 12q0-.95.463-2.15t1.612-2.288q1.15-1.087 3.075-1.825T12 5q2.925 0 4.85.738t3.075 1.824q1.15 1.088 1.613 2.288T22 12q0 2.9-2.05 4.95T15 19H9Zm.15-2h5.7q1.575 0 2.863-.838T19.6 14H4.4q.6 1.325 1.888 2.163T9.15 17ZM12 14Zm0-1Zm-8-1h16q0-.75-.4-1.625T18.262 8.75Q17.326 8 15.789 7.5T12 7q-2.25 0-3.775.5T5.762 8.75q-.937.75-1.35 1.625T4 12Zm8 0Z"/>`),
		g.Group(children),
	)
}