package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallpaperengine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.395 13.428a11.18 11.18 0 0 1 9.49 17.086m-9.49-17.085H9.865A4.356 4.356 0 0 0 5.5 17.796v20.342a4.354 4.354 0 0 0 4.364 4.365h20.342a4.356 4.356 0 0 0 4.368-4.365"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.006 13.429v-3.235h4.262v3.61m3.61 1.698l2.862-2.862l2.595 2.62l-2.862 2.862m2.099 7.847h3.236v-4.262h-3.61"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.498 18.097l2.862-2.862l-2.62-2.594l-2.862 2.861m-5.273 19.07a11.18 11.18 0 0 1-9.489-17.086"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.605 34.571h13.53a4.356 4.356 0 0 0 4.365-4.367V9.862a4.354 4.354 0 0 0-4.364-4.365H17.793a4.356 4.356 0 0 0-4.367 4.365"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.994 34.571v3.235h-4.262v-3.61m-3.61-1.698l-2.861 2.862l-2.595-2.621l2.862-2.862m-2.1-7.846h-3.235v4.262h3.61"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.502 29.903l-2.862 2.862l2.62 2.595l2.862-2.862"/><circle cx="24" cy="24" r="4.14" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}