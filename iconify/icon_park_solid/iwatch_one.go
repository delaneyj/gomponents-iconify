package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IwatchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSIwatchOne0"><g fill="none" stroke-width="4"><rect width="22" height="24" x="13" y="12" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 12v6"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M21 12h6"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 30v6"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M21 36h6"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M35 23.934L29 24"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M35 21v6"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m19 24l-6 .066"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M13 21v6m18-15V4H17v8m14 24v7a1 1 0 0 1-1 1H18a1 1 0 0 1-1-1v-7M7 20v8m34-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSIwatchOne0)"/>`),
		g.Group(children),
	)
}