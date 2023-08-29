package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Argenta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.797 16.203H13.321L4.5 23.182v12.649l8.821 6.98h12.476l8.821-6.98V23.182l-8.821-6.979zm8.821 12.467L16.746 42.811M29.91 19.457L6.857 37.696m10.811-21.493L4.5 26.622m31.387-12.623L43.5 7.976H28.067l-7.613 6.023h15.433zm-15.433 0v-8.81"/>`),
		g.Group(children),
	)
}