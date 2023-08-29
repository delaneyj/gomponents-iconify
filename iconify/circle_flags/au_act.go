package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AuAct(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsAuAct0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuAct0)"><path fill="#0052b4" d="M0 0h170l64 256l-64 256H0Z"/><path fill="#ffda44" d="M170 0h342v512H170z"/><path fill="#eee" d="m136 189l7 14.7l15.9-3.7l-7.2 14.7l12.7 10l-15.9 3.7v16.3l-12.6-10.2l-12.7 10.2v-16.3l-15.9-3.6l12.8-10.1L113 200l15.9 3.7zm-45-43l7 14.7l15.9-3.7l-7.2 14.7l12.7 10l-15.9 3.7v16.3l-12.6-10.2l-12.7 10.2v-16.3l-15.9-3.6l12.8-10.1L68 157l15.9 3.7zm19 92l5.6 17h17.8l-14.5 10.5l5.5 17l-14.5-10.5l-14.4 10.5l5.5-17L86.5 255h18zm-70-34l7 14.7l15.9-3.7l-7.2 14.7l12.7 10l-15.9 3.7v16.3l-12.6-10.2l-12.7 10.2v-16.3l-15.9-3.6l12.8-10.1L17 215l15.9 3.7zm51 99l-6 17l-17-8l8 17l-17 6l17 6l-8 17l17-8l6 17l6-17l17 8l-8-17l17-6l-17-6l8-17l-17 8z"/><g transform="translate(0 22)"><path fill="#0052b4" d="M276 200a17 17 0 0 0-17 17h-51a17 17 0 0 0 17 17a17 17 0 0 0 17 17a17 17 0 0 0 17 17h71v-68z"/><path fill="#eee" d="M384 200a17 17 0 0 1 17 17h51a17 17 0 0 1-17 17a17 17 0 0 1-17 17a17 17 0 0 1-17 17h-71v-68z"/><path fill="#0052b4" d="M285.5 200v68c0 34 44.5 44 44.5 44s44.5-10 44.5-44v-68z"/><path fill="#eee" d="m294.4 204l-2.5 7.5l26.8 10l-26.8 10l2.5 7.5l35.6-13.3l35.7 13.3l2.5-7.5l-26.9-10l26.9-10l-2.5-7.5l-35.7 13.3zm25 31.2V254h-6.2v-9.2h-15V284h63.7v-38.8h-15v8.8h-7.5v-18.8z"/><circle cx="330" cy="297.2" r="9" fill="#eee"/></g></g>`),
		g.Group(children),
	)
}