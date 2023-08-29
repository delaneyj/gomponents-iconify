package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trenitalia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.285 13.264c-10.904.084-12.708 1.887-15.14 6.333h12.037ZM30.754 28.865a9.723 9.723 0 0 1-7.886 7.885a29.294 29.294 0 0 1-3.984.419C17.29 37.21 4.5 37.21 4.5 37.21h0s4.865-.042 5.536-.042a43.691 43.691 0 0 0 4.739-.335c3.34-.282 6.687-3.764 7.547-7.003l.965-3.271a5.93 5.93 0 0 1 5.787-4.781l9.405-.229h0l-3.281.08a4.317 4.317 0 0 0-4.087 3.799Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.89 22.532l-8.471.838l-5.703 11.324c4.53 0 6.836-.418 9.478-4.193a77.617 77.617 0 0 0 4.696-7.969ZM43.5 10.79H25.638c-6.04 0-9.145 2.055-11.032 5.83l-1.468 2.976h8.134c2.854-5.535 7.048-8.678 15.645-8.806Z"/>`),
		g.Group(children),
	)
}