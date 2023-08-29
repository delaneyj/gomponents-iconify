package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.5 9.5c0 .918-.247 1.778-.68 2.518m-1.424 1.558A5 5 0 0 1 5.5 9.5m2.196-4.955a3.001 3.001 0 0 1 2.693-1.93l.111-.004a2.893 2.893 0 0 1 3 2.891V9.5c0 .388-.074.759-.208 1.099"/><path d="M11.957 12.123A3 3 0 0 1 7.5 9.5v-2m-4-4l14 14m-7-3v4"/></g>`),
		g.Group(children),
	)
}