package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vneid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.824 4.5h26.352c1.844 6.851 2.108 11.068 1.792 16.338c-1.265 10.54-4.164 13.176-7.59 17.128C29.534 40.338 25.318 43.5 24 43.5s-5.534-3.162-7.378-5.534c-3.426-3.953-6.325-6.588-7.59-17.128c-.316-5.27-.052-9.487 1.792-16.338Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.713 16.095v7.378l-4.848-7.378v7.378v-7.378l4.85 7.326l-.002-7.326Zm5.29 6.463a1.812 1.812 0 0 1-1.577.923a1.822 1.822 0 0 1-1.815-1.83v-1.188c0-1.01.813-1.83 1.815-1.83s1.815.82 1.815 1.83v.594h-3.63m5.524-4.962v7.378m1.897 0v-7.378h1.647c1.768 0 3.202 1.445 3.202 3.228v.922c0 1.783-1.434 3.228-3.202 3.228h-1.647ZM11.12 16.095l2.424 7.378l2.424-7.378"/>`),
		g.Group(children),
	)
}