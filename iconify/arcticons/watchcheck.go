package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watchcheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="18.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.019" cy="24.012" r="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M24 5.5c.52-.235.973-1.013 1.29-1.49a3.098 3.098 0 0 0 .51-1.51M42.5 24c.235.52 1.013.973 1.49 1.29a3.098 3.098 0 0 0 1.51.51M24 42.5c-.52.235-.973 1.013-1.29 1.49a3.098 3.098 0 0 0-.51 1.51M5.5 24c-.235-.52-1.013-.973-1.49-1.29a3.098 3.098 0 0 0-1.51-.51M14.75 7.979a4.205 4.205 0 0 0 .373-1.935a3.098 3.098 0 0 0-.314-1.563M40.022 14.75a4.205 4.205 0 0 0 1.934.373a3.098 3.098 0 0 0 1.563-.314M33.25 40.022a4.205 4.205 0 0 0-.373 1.934a3.098 3.098 0 0 0 .314 1.563M7.979 33.25a4.205 4.205 0 0 0-1.935-.373a3.098 3.098 0 0 0-1.563.314M7.979 14.75a4.204 4.204 0 0 0-.645-1.862a3.098 3.098 0 0 0-1.053-1.197M33.25 7.979a4.204 4.204 0 0 0 1.862-.645a3.098 3.098 0 0 0 1.197-1.053m3.713 26.969a4.204 4.204 0 0 0 .644 1.862a3.098 3.098 0 0 0 1.054 1.197m-26.97 3.713a4.204 4.204 0 0 0-1.862.644a3.098 3.098 0 0 0-1.197 1.054M24 19.506v-14M28.51 24h14m-18.491 4.512v14M19.51 24h-14"/>`),
		g.Group(children),
	)
}