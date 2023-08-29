package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cred(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.013 33.553L24.027 43.5l-17.04-9.947V4.5h34.026v29.053z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.651 21.872v9.232L24.02 38.5l-12.671-7.396V14.5H24m-8.355-5h21.006v7.715"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.485 29.71l-6.472 3.79l-8.302-4.846v-6.056"/>`),
		g.Group(children),
	)
}