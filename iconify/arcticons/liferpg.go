package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liferpg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="33.545" cy="17.232" r="1.165" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.36" cy="20.044" r="1.165" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.731" cy="20.044" r="1.165" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.545" cy="22.855" r="1.165" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.158 36.107c-2.665 3.38-5.7-2.33-5.658-4.801c.1-5.793 2.974-13.45 4.903-16.789a5.675 5.675 0 0 1 2.258-2.206a11.8 11.8 0 0 1 4.95-1.458c.567-.07 1.615 1.417 1.615 1.417h11.547s1.049-1.487 1.616-1.417a11.8 11.8 0 0 1 4.95 1.458a5.675 5.675 0 0 1 2.258 2.206c1.93 3.338 4.803 10.996 4.903 16.789c.042 2.472-2.993 8.181-5.658 4.801c-2.115-2.683-4.12-7.421-7.14-7.933a48.464 48.464 0 0 0-6.702-.35a50.274 50.274 0 0 0-6.702.35c-3.08.402-5.024 5.25-7.14 7.933Z"/><circle cx="24" cy="16.858" r="1.703" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.632 22.995v-1.563m-3.05-1.484h1.565m1.485-3.046v1.563m3.048 1.483h-1.564"/>`),
		g.Group(children),
	)
}