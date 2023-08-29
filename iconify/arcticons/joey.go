package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Joey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.666 14.981a6.075 6.075 0 0 0-9.326-3.609c2.297-.055 2.629 1.26 2.629 1.26s-6.337-.457-5.369 5.41a1.897 1.897 0 0 1 1.938-.802s-3.672 4.9 3.297 5.452"/><circle cx="16.86" cy="21.999" r="2.269" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 37.718c8.1 0 14.667-4.484 14.667-11.762c0-6.89-6.567-12.702-14.667-12.702S9.333 19.066 9.333 25.956c0 7.278 6.567 11.762 14.667 11.762Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.092 32.57S27.819 34.397 24 34.397s-5.092-1.827-5.092-1.827M24 30.8a2.783 2.783 0 0 0 3.044-2.796c0-1.218-.498-3.957-3.044-3.957s-3.044 2.74-3.044 3.957A2.783 2.783 0 0 0 24 30.8Zm7.334-15.819a6.075 6.075 0 0 1 9.326-3.609c-2.297-.055-2.629 1.26-2.629 1.26s6.337-.457 5.369 5.41a1.897 1.897 0 0 0-1.938-.802s3.672 4.9-3.297 5.453"/><circle cx="31.14" cy="21.999" r="2.269" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}