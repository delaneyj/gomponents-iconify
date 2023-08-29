package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeatGame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="33.475" cy="12.483" r="6.983" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.401 16.007s-3.63 2.892-2.085 4.329c1.213 1.127.786.625 2.472-.225c0 0-.42 2.362 1.016 2.368c0 0-2.543 1.015-.872 2.365c.956.773 1.999 1.417 3.762-1.145"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M39.525 16.717s4.053 1.388 2.207 3.525c-1.142 1.322-2.916-.572-2.916-.572c-.319.013.936 3.452-2.075 4.337c-.974 1.743-.753 2.736-2.328 2.466c-1.697-.29-.868-2.791-.868-2.791a4.917 4.917 0 0 0-1.85.018"/><circle cx="34.275" cy="12.796" r="2.199" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M32.684 16.855h2.89"/><circle cx="14.437" cy="40.366" r="2.134" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M16.571 39.957V28.983m8.106-16.808L12.935 23.917"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.94 28.104l-5.803 5.803"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="m24.67 16.477l-14 13.998m8.597 7.483l-1.271 1.272M30.1 27.126l-3.133 3.133m-3.396-.172L18.39 35.27m-3.803-4.5l-3.82 3.82m12.434-12.433l-4.296 4.297m-.391 4.567l1.362-1.362"/><circle cx="14.437" cy="40.366" r="2.134" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M16.571 39.957V28.983"/><circle cx="23.279" cy="37.164" r="2.134" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M25.413 36.755V25.781m-8.842 3.202l8.842-3.202"/><circle cx="7.478" cy="35.662" r="1.692" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.17 35.553V24.58s2.612.602 2.651 2.331"/>`),
		g.Group(children),
	)
}