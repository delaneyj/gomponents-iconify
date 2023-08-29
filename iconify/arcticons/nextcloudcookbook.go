package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nextcloudcookbook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.45v35.1a2 2 0 0 0 1.95 2h2.38V4.5h-2.38A2 2 0 0 0 8.4 6.45Zm4.33-1.95v39h24.92a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95Zm9.06 13.04v-4.29m-2.25 4.49v-4.49m4.36 4.29v-4.29m2.27 4.49v-4.49"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.17 17.74a3.32 3.32 0 0 1-6.63 0m3.31 17.01V21.11m9.93 13.64v-21.5c-5 1.87-4.75 11.25 0 13.6"/>`),
		g.Group(children),
	)
}