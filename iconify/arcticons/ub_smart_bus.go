package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UbSmartBus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a16.246 16.246 0 0 1 16.246 16.246C40.246 34.73 24 43.5 24 43.5S7.754 34.73 7.754 20.746A16.246 16.246 0 0 1 24 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 9.384a17.218 17.218 0 0 1 2.998.25a13.628 13.628 0 0 1 3.391 1.04a3.311 3.311 0 0 1 1.53 1.142a3.22 3.22 0 0 1 .36 1.205a53.829 53.829 0 0 1 .857 6.143v9.638H31.6v1.537a1.313 1.313 0 1 1-2.62 0v-1.536h-9.96v1.536a1.313 1.313 0 1 1-2.62 0v-1.536h-1.537v-9.639a55.127 55.127 0 0 1 .84-6.143a3.22 3.22 0 0 1 .36-1.205a3.312 3.312 0 0 1 1.53-1.142a13.63 13.63 0 0 1 3.391-1.04A17.22 17.22 0 0 1 24 9.385Zm-6.28 13.88a1.376 1.376 0 1 0 1.37 1.382v-.012a1.37 1.37 0 0 0-1.37-1.37Zm12.56 0a1.376 1.376 0 1 0 1.377 1.376v-.006a1.365 1.365 0 0 0-1.36-1.37ZM16.4 28.802h15.2m-16.736-9.638h18.272M20.972 25.64h6.056m-6.056-2h6.056"/>`),
		g.Group(children),
	)
}