package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iconic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.139 22.246C15.84 9.692 36.076 9.8 36.076 24.096S15.84 38.503 19.14 25.949"/><circle cx="23.618" cy="17.692" r="1.422" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.198" cy="18.919" r="1.422" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.644 22.672a1.422 1.422 0 0 0-1.411 1.272c.3.259.594.518.882.757a5.216 5.216 0 0 0-.669.16a1.422 1.422 0 1 0 1.198-2.19Z"/><circle cx="23.618" cy="30.494" r="1.422" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.198" cy="29.268" r="1.422" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.471 25.4v-2.577c-7.536-1.208-14.446.38-14.446 1.289s6.91 2.497 14.446 1.289Zm-2.736-2.902v3.228m2.736-.326c1.815 1.619 3.427-.332 5.643-.698c-1.669-1.39-3.539-3.487-5.643-1.879M11.129 4.5a3.295 3.295 0 0 0-3.302 3.298h32.345A3.295 3.295 0 0 0 36.87 4.5ZM7.827 40.202a3.295 3.295 0 0 0 3.302 3.298H36.87a3.295 3.295 0 0 0 3.302-3.298Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.043 11.057L7.827 24m32.346 0l2.784-12.944c.342-1.616-1.202-3.01-2.784-3.258H7.827c-1.582.248-3.126 1.642-2.784 3.259m0 25.887L7.827 24m32.346 0l2.784 12.944c.342 1.616-1.202 3.01-2.784 3.258H7.827c-1.582-.248-3.126-1.642-2.784-3.258"/>`),
		g.Group(children),
	)
}