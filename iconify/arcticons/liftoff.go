package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liftoff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M23.968 43.5c.482 0 4.355-2.488 4.355-6.953h-8.71c0 4.465 3.872 6.953 4.355 6.953Zm-8.306-23.328c-.717 0-6.474 4.653-6.474 7.292v10.5c1.5-1.5 6.474-4.066 6.474-4.066V20.172Zm16.676 0c.717 0 6.474 4.653 6.474 7.292v10.5c-1.5-1.5-6.474-4.066-6.474-4.066V20.172Z"/><path d="M29.943 10.152C28.423 7.912 26.473 5.663 24 4.5c-.924 0-8.338 5.465-8.338 13.312v16.086c5.675 0 10.975.15 16.676 0V17.812"/><circle cx="23.769" cy="19.13" r="4.137"/></g>`),
		g.Group(children),
	)
}