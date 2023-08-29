package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeWorkouts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.077 12.27s-.611-.321-1.008-2.14s3.086-5.18 3.086-5.18s1.07-.519 5.5-.442s5.486.794 6.403 1.894C32.055 11.2 39.563 20.75 41.345 28.28c.406 1.122.966.51.966 2.321c0 1.298-.71 1.656-.9 1.868c-4.614 5.164-15.92 6.875-20.963 6.875s-5.011-1.161-6.417-1.161s-2.964 4.645-6.44 5.317c-2.59-5.195-2.483-15.707 0-21.26c5.844 0 6.776 2.245 9.19 2.245c2.063 0 1.742-1.283 4.997-1.283s7.746 3.712 7.746 3.712c-2.544-3.162.292-10.868-5.202-15.997c-1.032-.962-1.513 2.2-4.722 2.2c-.473.28-2.175.503-2.175.503"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.884 28.224c.894 1.02 2.731 2.403 4.439 2.38m-22.85-1.069c-1.971 2.368-2.231 4.935-4.217 6.005m7.257-4.935c2.965 3.01 8.328 2.81 11.46-.398m.133-21.467c-.137.573-.87 1.455-1.41 1.604M18.73 6.448c-.401.481-2.35 2.957-2.304 3.358s-.24 2.406.367 2.624s2.67-.15 2.647-1.02s-.195-1.318-.149-1.742s1.421-2.52 1.421-2.52"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.747 6.081c-.504.608-2.166 2.957-2.143 3.942s.126 1.708.355 1.926s1.468-.295 1.468-.295"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.87 12.452s.496.54.553 1.124s-.756 1.146-1.5 1.352s-1.789-1.192-1.789-1.719s.825-1.26.825-1.26m6.619-2.279c.026.741.207 1.544-.432 1.777s-1.745.176-1.745.176"/>`),
		g.Group(children),
	)
}