package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 4a1 1 0 0 1 1 1v1.016c9.157.501 16.494 7.838 16.995 16.995c.016.293-.34.343-.558.147a7.182 7.182 0 0 0-4.822-1.85a7.192 7.192 0 0 0-5.433 2.47a.663.663 0 0 1-.497.233a.656.656 0 0 1-.453-.187A8.981 8.981 0 0 0 25 21.364V39.5a2.5 2.5 0 0 0 5 0V38a1 1 0 0 1 2 0v1.5a4.5 4.5 0 1 1-9 0V21.36a8.981 8.981 0 0 0-5.272 2.474a.62.62 0 0 1-.429.177a.626.626 0 0 1-.47-.22a7.192 7.192 0 0 0-5.444-2.483c-1.854 0-3.545.7-4.822 1.85c-.218.196-.574.146-.557-.147c.5-9.157 7.837-16.494 16.994-16.995V5a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}