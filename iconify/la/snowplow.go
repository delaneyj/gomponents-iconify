package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowplow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.5 6a3.005 3.005 0 0 0-2.875 2.188l-1.313 4.53l-2-.655l-.625 1.874l1.47.5l-.282.282A2.931 2.931 0 0 0 4 16.813V19h-.813l-.156.813l-1 5L1.781 26H30.22l-.25-1.188l-1-5l-.157-.812H28v-2.188c0-.796-.313-1.53-.875-2.093l-.281-.281l1.468-.5l-.625-1.876l-2 .657l-1.312-4.531A3.005 3.005 0 0 0 21.5 6zm0 2h11a.96.96 0 0 1 .938.719L23.655 13H8.344l1.219-4.281A.96.96 0 0 1 10.5 8zm-3.063 7h17.125l1.157 1.125a.944.944 0 0 1 .281.688V19h-2v-2h-4v2h-8v-2H8v2H6v-2.188c0-.265.094-.5.281-.687zm-2.625 6h22.375l.594 3H4.22z"/>`),
		g.Group(children),
	)
}