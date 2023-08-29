package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pulseaudiortpreceiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="34.1" cy="21.8" r="1.6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26.8" cy="24.1" r="1.7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.2" cy="25.3" r="1.7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.4" cy="23.2" r=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="22.6" cy="15.5" r=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.6" cy="31.1" r=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="25.4" cy="34.5" r=".5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14.9" cy="36.1" r=".4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.3 16.8v4m-2.3-2h4.4m10.1 5.4V18l-5.6 1.2m0 6.2v-6.3m13.6 13.1L37 30a1.78 1.78 0 0 0-1.2-1.9l-2.4-.5a1.36 1.36 0 0 0-1.7 1.3L29.6 40c-.3 1.1-.1 1.8.8 1.8l2.2.7a1.68 1.68 0 0 0 2.1-1l1.3-6.6m-1.8-4.4l-1 5.6m-.6 2.7l-.1.7m-17.1-.9l.1.9m-1.1-5.9l-.6-3.3M35 40.6c.2 0 3.4.7 3.4.7l2-1.9l1.2-6l-1.1-2.2l-3.5-.9m-25.9.1l-3.9.9l-.9 1.9L7.6 40l1.8 1.3l3.7-.7m-3.4-1.7l.5 2.2m27.8.1l3-15.3M9.2 36.3L7.1 25.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.7 24.3c.5-5.9-4-13.2-16.2-12.9c-11.6.1-15.8 7.2-15.2 12.9c.2 1.8-2.9 2.4-2.7 0C7.1-1.6 42.1 0 42.4 24.6a1.37 1.37 0 1 1-2.7-.3Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.7 22.7c.2-5.6-3.3-10.3-8.7-12.9M28.5 9c-11-2.4-20 2.8-20.2 13.3m10 17.8c.1 1.1-.3 1.7-1 1.7l-2.1.5c-2-.2-1.6-.8-1.9-1.3l-2.1-11c-.1-1.2.4-1.7 1-1.9l2-.4c1.1-.1 1.9.2 2.1 1.3Z"/>`),
		g.Group(children),
	)
}