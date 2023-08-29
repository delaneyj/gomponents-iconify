package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tungngern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.887 26.925c1.293 1.098 2.074.908 4.38 2.88c6.179 5.283 4.235 12.412-10.136 12.695c-14.372-.283-16.86-7.453-10.682-12.736c2.306-1.972 3.512-1.745 4.926-2.839"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.374 24.657h11.513v2.269H15.374z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.466 24.658l-3.784-5.035c1.732-2.574 4.17-1.33 6.556 0c2.335 1.3 4.49 3.04 7.59 2.016c2.846-.94 3.41-2.036 5.068-3.058l-4.012 5.765"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.846 18.794v-5.898h8.258v8.733"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.091 19.439a3.005 3.005 0 0 1 5.544 2.318m2.305-.303c.148-1.505.988-3.634 2.197-6.289m-.624-4.397l-5.1-1.102l-1.298 3.083m-4.947 5.74c-.615-2.643-1.903-5.735-4.372-8.761l8.02-1.657c.66 1.265 1.419 2.82 1.612 3.764"/><circle cx="34.116" cy="12.105" r="6.605" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.911 12.308c.945.319 1.464 1.374 1.154 2.345c-.31.97-1.336 1.504-2.28 1.186l-2.834-.956l2.254-7.062l2.834.956c.944.319 1.463 1.374 1.153 2.345c-.31.97-1.336 1.505-2.28 1.186Zm.021.02l-2.858-.957m3.364-2.797l.281-.882m-2.818 8.827l.282-.883"/>`),
		g.Group(children),
	)
}