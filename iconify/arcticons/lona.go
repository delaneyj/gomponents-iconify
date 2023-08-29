package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lona(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.654 23.993a15.813 15.813 0 1 1-31.625.027v-.027a7.882 7.882 0 0 0-4.36-7.074a3.95 3.95 0 0 1 3.553-7.055a15.742 15.742 0 0 1 8.708 14.13a7.912 7.912 0 1 0 7.912-7.913h0a3.95 3.95 0 0 1 0-7.9a15.809 15.809 0 0 1 13.067 6.916"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.542 17.637c-1.797-1.727-.272-4.765.704-5.506m1.684 8.376c-.918-2.714 3.434-6.246 6.526-5.232c1.296.425 3.161 4.31 3.038 8.128a2.496 2.496 0 0 1-.84 1.626"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.73 25.029c.581-2.967 6.93-2.638 7.925-1.036m-9.025 4.054c1.926-1.378 4.646.744 4.992 2.217m-7.923.634a3.458 3.458 0 0 1 3.517 3.37m-8.155-2.401c2.11.926 1.925 3.578 1.697 4.372m-6.24-6.405a3.032 3.032 0 0 1-.383 4.27m-1.798-7.611a3.248 3.248 0 0 1-.884 2.904m-.04-9.527a3.136 3.136 0 0 1-.88 2.917"/><circle cx="7.594" cy="14.316" r=".75" fill="currentColor"/><circle cx="9.2" cy="12.073" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}