package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Projectm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.631 13.688v7.222m-.001-4.243a1.81 1.81 0 0 0 1.806 1.806h0a1.81 1.81 0 0 0 1.806-1.806v-1.173a1.81 1.81 0 0 0-1.806-1.806h0a1.81 1.81 0 0 0-1.805 1.806M35.59 17.57a1.747 1.747 0 0 1-1.535.903h0a1.81 1.81 0 0 1-1.805-1.806v-1.173a1.81 1.81 0 0 1 1.806-1.806h0a1.747 1.747 0 0 1 1.534.903m-5.833 2.979a1.747 1.747 0 0 1-1.535.903h0a1.81 1.81 0 0 1-1.806-1.806v-1.173a1.81 1.81 0 0 1 1.806-1.806h0a1.81 1.81 0 0 1 1.806 1.806v.632h-3.611"/><path fill="currentColor" d="M24.931 11.522a.75.75 0 0 1-.75.75a.75.75 0 0 1-.75-.75a.75.75 0 0 1 .75-.75a.75.75 0 0 1 .75.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.181 13.688v5.417a1.81 1.81 0 0 1-1.805 1.805h0a1.782 1.782 0 0 1-1.264-.542m16.361-6.68h1.896m-.903-1.534v6.319m-23.78-4.785v4.785m0-2.979a1.81 1.81 0 0 1 1.805-1.806h0m3.712 4.785h0a1.81 1.81 0 0 1-1.806-1.806v-1.173a1.81 1.81 0 0 1 1.806-1.806h0a1.81 1.81 0 0 1 1.805 1.806v1.173a1.81 1.81 0 0 1-1.805 1.806Zm-2.458 18.755V24.717L24 37.228l6.255-12.511v12.511"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}