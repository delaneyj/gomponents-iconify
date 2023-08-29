package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="24" r="21.5"/><path d="m29.868 44.689l3.665-7.568l-3.449-4.632H17.916l-3.449 4.632l3.665 7.568"/><path d="m45.489 24.812l-6.064-5.824l-5.581 1.929l-3.76 11.572l3.449 4.632l8.33-1.147"/><path d="M31.413 3.813L24 7.781v5.984l9.844 7.152l5.581-1.929l1.483-8.277m-33.816.001l1.483 8.276l5.581 1.929L24 13.765V7.781l-7.414-3.968M6.137 35.974l8.33 1.147l3.449-4.632l-3.76-11.572l-5.581-1.929l-6.065 5.825"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.18 22.883c.24-1.352 1.653-2.419 3.006-2.15c.888.177 1.482.972 1.41 1.904c-.055.692-.394 1.376-.947 1.796c-1.024.778-4.249 2.875-4.249 2.875h4.384m1.631-.558c.389.383.851.558 1.959.558h.671c.914 0 1.785-.74 1.946-1.653h0c.16-.914-.45-1.654-1.363-1.654m-2.243-2.757c.525-.381 1.049-.555 2.156-.552l.67.001c.914 0 1.524.74 1.363 1.654h0c-.161.913-1.032 1.654-1.946 1.654m-1.684 0h1.684"/>`),
		g.Group(children),
	)
}