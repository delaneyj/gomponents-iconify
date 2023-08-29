package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VillageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.445 6.168a1 1 0 0 1 1.11 0l6 4A1 1 0 0 1 24 11v8a1 1 0 0 1-1 1H11a1 1 0 0 1-1-1v-8a1 1 0 0 1 .445-.832l6-4ZM16 18h2v-4h-2v4Zm4 0v-5a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v5h-2v-6.465l5-3.333l5 3.333V18h-2Zm14.496-5.868a1 1 0 0 0-.992 0l-7 4A1 1 0 0 0 26 17v10a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V17a1 1 0 0 0-.504-.868l-7-4ZM37 26h3v-8.42l-6-3.428l-6 3.428V26h3v-6a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v6Zm-2 0v-5h-2v5h2Zm-11.553 2.106l-8-4a1 1 0 0 0-.894 0l-8 4A1 1 0 0 0 6 29v12a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V29a1 1 0 0 0-.553-.894ZM16 34v6h-2v-6h2Zm2-1v7h4V29.618l-7-3.5l-7 3.5V40h4v-7a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}