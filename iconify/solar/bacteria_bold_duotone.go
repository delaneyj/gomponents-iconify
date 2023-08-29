package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BacteriaBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 15a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M14.465 7.171s1.243-.171 2.121.707c.879.879.707 2.122.707 2.122M9 9.293s-1.243.171-2.121-.707C6 7.707 6.172 6.464 6.172 6.464M19 13.136s-1.162.473-1.483 1.673c-.322 1.2.448 2.19.448 2.19m-4.545.773L15 19"/><path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2Zm-6.393 8.787a.75.75 0 0 0-1.386-.574l-.414 1a.75.75 0 0 0 1.386.574l.414-1Zm8.893.463a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 1 .75-.75Zm-2.5-7h-2a.75.75 0 0 0 0 1.5h2a.75.75 0 0 0 0-1.5Z" clip-rule="evenodd" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m12.5 9.5l-1-1"/></g>`),
		g.Group(children),
	)
}