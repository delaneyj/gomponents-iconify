package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stickerly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.83 45.268C10.023 43.67 1.835 34.032 2.543 22.628c.664-10.707 9.378-19.42 20.085-20.085c11.404-.708 21.043 7.48 22.64 18.287a5.996 5.996 0 0 1-1.704 5.106L25.936 43.564a5.996 5.996 0 0 1-5.106 1.704Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="17.402" cy="18.3" rx="3.046" ry="4.3"/><ellipse cx="30.598" cy="18.3" rx="3.046" ry="4.3"/><path d="M33.644 27.834C32.184 32.454 27.982 34 24.39 34c-3.593 0-7.256-1.322-9.255-6.165"/></g>`),
		g.Group(children),
	)
}