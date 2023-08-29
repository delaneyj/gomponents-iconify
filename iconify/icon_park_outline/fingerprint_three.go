package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FingerprintThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42.43 12.04C37.703 7.382 31.155 4.5 23.919 4.5c-7.193 0-13.706 2.849-18.428 7.457m1.176 17.517v-.057c0-9.573 7.76-17.334 17.333-17.334m7.17 1.547c5.995 2.728 10.163 8.77 10.163 15.787v.013M14.25 37v-7.584c0-5.384 4.365-9.75 9.75-9.75s9.75 4.366 9.75 9.75V37"/><path d="M17.526 43.5c1.963-.485 3.224-2.554 3.224-4.304V30.5a3.25 3.25 0 0 1 6.5 0v8.696"/></g>`),
		g.Group(children),
	)
}