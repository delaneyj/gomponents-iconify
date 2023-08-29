package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyBoy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="5.333" d="M28.81 11.707c3.158-.846 3.502.53 5.232.784a14.025 14.025 0 0 0 3.92.004c1.388-.112 1.509.208 2.132 1.755a105.189 105.189 0 0 1 3.32 10.28c.788 3.244-3.631 3.932-8.333 7.367c-1.493 1.091-4.735 2.067-8.437 3.059s-6.999 1.767-8.837 1.568c-5.79-.625-9.961.988-10.902-2.215a105.19 105.19 0 0 1-2.267-10.562c-.234-1.651-.29-1.99.969-2.585A14.024 14.024 0 0 0 9 19.198c1.371-1.084.98-2.448 4.139-3.294l7.835-2.099Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m31.586 25.939l-2.233-8.316l-13.73 3.672l2.232 8.315l6.866-1.836Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M30.607 15.543c1.45.58 3.573 8.309 4.027 11.144s-3.28 4.142-8.706 5.595s-9.314 2.189-10.339-.494s-3.052-10.437-2.086-11.663s5.498-2.243 8.36-3.01s7.293-2.151 8.743-1.572Z"/><circle cx="35.865" cy="21.042" r="1.027" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="rotate(-10.012 35.866 21.042)"/><circle cx="38.519" cy="17.827" r="1.027" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="rotate(-10.012 38.519 17.827)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.021 24.597l1.278 5.301m2.012-3.29L7.01 27.886"/>`),
		g.Group(children),
	)
}