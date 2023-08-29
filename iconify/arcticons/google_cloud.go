package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.565 11.514c-6.635-5.857-16.76-5.225-22.617 1.41c-1.627 1.843-2.807 4.125-3.451 6.497"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.417 17.611c3.172-3.483 8.483-4.058 12.145-1.094M23.92 40.496h8.907c6.397.05 11.623-5.095 11.673-11.492a11.583 11.583 0 0 0-5.094-9.685a16.061 16.061 0 0 0-4.84-7.806l-5.004 5.004a8.894 8.894 0 0 1 3.265 7.055v.888a4.453 4.453 0 1 1 0 8.907H23.92M8.022 38.17a11.52 11.52 0 0 0 6.992 2.326h8.906v-7.13h-8.906a4.428 4.428 0 0 1-1.839-.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.014 17.368C8.617 17.406 3.462 22.623 3.5 29.02c.021 3.572 1.69 6.972 4.522 9.15l5.153-5.204c-2.241-1.012-3.225-3.65-2.212-5.892a4.453 4.453 0 0 1 8.116 0l5.167-5.166a11.57 11.57 0 0 0-9.232-4.54Z"/>`),
		g.Group(children),
	)
}