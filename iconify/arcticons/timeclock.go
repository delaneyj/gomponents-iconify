package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timeclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.434 33.099H14.566m0 10.401h18.868l-2.968-10.401H16.898L14.566 43.5z"/><circle cx="24" cy="18.677" r="10.388" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 18.677v-7.364m0 7.364l-3.85 1.861M15.27 40.36H13a3.177 3.177 0 0 1-3.177-3.177h0V7.677A3.177 3.177 0 0 1 12.999 4.5h22.002a3.177 3.177 0 0 1 3.176 3.176h0v29.507a3.177 3.177 0 0 1-3.176 3.177h-2.463"/>`),
		g.Group(children),
	)
}