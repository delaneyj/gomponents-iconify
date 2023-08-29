package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lung(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4.571"><path d="M18.244 22.263c-.292-4.955-1.467-7.521-3.523-7.698c-3.32-.284-8.41 5.144-9.92 12.772c-1.509 7.628-.764 13.092 1.211 13.46c1.975.37 4.844-1.78 6.205-1.78c1.361 0 6.079 1.04 6.079-.844V30.07m11.328-7.7c.281-5.025 1.458-7.627 3.53-7.805c3.32-.284 8.41 5.144 9.919 12.772c1.509 7.628.764 13.092-1.21 13.46c-1.975.37-4.845-1.78-6.206-1.78c-1.36 0-6.079 1.04-6.079-.844V30.07"/><path d="M20.556 5v14.91c-.044 2.292-1.956 3.438-5.737 3.438M27.292 5v14.91c.043 2.292 1.956 3.438 5.736 3.438"/><path stroke-linejoin="round" d="M15 30.07c3.955 0 6.955-.939 9-2.818c2.06 1.88 5.069 2.818 9.029 2.818"/></g>`),
		g.Group(children),
	)
}