package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Interview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#FCEA2B"><ellipse cx="20" cy="26.021" stroke="#FCEA2B" stroke-miterlimit="10" stroke-width="2" rx="13.8" ry="13.82"/><path d="M18.398 39.91C15.244 44.248 11 46 5 46c2.874-1.916 5.748-5.668 6.864-8.617l6.534 2.527z"/><circle cx="52" cy="41.28" r="13.8" stroke="#FCEA2B" stroke-miterlimit="10" stroke-width="2"/><path d="M53.602 55.19C56.756 59.528 61 61.28 67 61.28c-2.874-1.916-5.748-5.668-6.864-8.617l-6.534 2.527z"/></g><circle cx="13" cy="26" r="2"/><circle cx="20" cy="26" r="2"/><circle cx="27" cy="26" r="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M11.864 37.383C10.748 40.333 7.874 44.084 5 46c6 0 10.244-1.752 13.398-6.09"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2" d="M22.275 39.816C28.925 38.73 34 32.957 34 26c0-7.732-6.268-14-14-14S6 18.268 6 26c0 3.334 1.165 6.395 3.11 8.8"/><circle cx="59" cy="41.28" r="2"/><circle cx="52" cy="41.28" r="2"/><circle cx="45" cy="41.28" r="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M60.136 52.663c1.116 2.95 3.99 6.701 6.864 8.617c-6 0-10.244-1.752-13.398-6.09"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2" d="M49.725 55.096C43.075 54.01 38 48.237 38 41.28c0-7.732 6.268-14 14-14s14 6.268 14 14c0 3.334-1.165 6.395-3.11 8.8"/>`),
		g.Group(children),
	)
}