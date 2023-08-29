package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmptyNest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#a57939" d="M11.59 31.688c0-6.796 10.929-12.305 24.41-12.305s24.41 5.51 24.41 12.305v1.056c0 5.691-2.009 12.49-5.062 16.344c-3.089 3.898-9.593 6.154-19.348 6.154s-15.854-1.826-19.347-6.154c-3.088-3.826-5.063-9.68-5.063-16.344Z"/><path fill="#6a462f" d="M56.616 47.215c0 2.108-7.333 8.027-17.63 8.027S20.34 53.533 20.34 51.425s7.458.644 19.39-1.374c10.154-1.717 16.886-4.944 16.886-2.836Z"/><ellipse cx="36" cy="33.754" fill="#6a462f" rx="17.341" ry="7.084"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23.156 36.807c0-1.692 5.773-4.063 12.844-4.063c7.016 0 12.844 2.371 12.844 4.064"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M52.324 33.688c0 3.891-7.308 7.045-16.324 7.045s-16.324-3.154-16.324-7.045m33.625 8.841c-1.584 2.24-5.687 4.008-10.885 4.686M36 24.643a28.57 28.57 0 0 1 12.474 2.5"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.59 31.688c0-6.796 10.929-12.305 24.41-12.305s24.41 5.51 24.41 12.305v1.056c0 5.691-2.009 12.49-5.062 16.344c-3.089 3.898-9.593 6.154-19.348 6.154s-15.854-1.826-19.347-6.154c-3.088-3.826-5.063-9.68-5.063-16.344Z"/><ellipse cx="36.188" cy="36.923" rx="13.031" ry="4.179"/>`),
		g.Group(children),
	)
}