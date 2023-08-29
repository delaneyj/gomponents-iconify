package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VietnamAirlines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.156 30.01c-1.875-.502-2.72-2.266-.237-4.748c1.117-1.118-.765-10.692-4.105-12.89"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.814 12.371c.85 3.024-1.675 6.091-2.238 10.516c-.564 4.424 1.704 6.953 4.206 8.412"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.576 22.887c-3.105-2.098-6.33-3.557-9.973-2.782c3.22 2.714 3.903 7.227 5.699 9.023c1.564 1.564 4.981 3.46 6.716 3.46"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.302 29.128c-4.624.136-6.4 2.79-7.802 4.545c2.92-.782 7.697 1.046 10.38 1.764s6.335-.705 9.12-2.312m11.698-3.997c4.624.136 6.4 2.79 7.802 4.545c-2.92-.782-7.697 1.046-10.38 1.764s-6.335-.705-9.12-2.312m5.186-20.754c-.85 3.024 1.675 6.091 2.238 10.516c.564 4.424-1.704 6.953-4.206 8.412"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.424 22.887c3.105-2.098 6.33-3.557 9.973-2.782c-3.22 2.714-3.902 7.227-5.699 9.023c-1.564 1.564-4.981 3.46-6.716 3.46m-4.138-2.578c1.87-.5 2.72-2.266.237-4.748c-1.117-1.118.765-10.692 4.105-12.89"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.394 28.316c0-1.944 5.21-1.944 5.21 0"/>`),
		g.Group(children),
	)
}