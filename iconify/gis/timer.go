package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M42 0a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h3v5.295C23.364 15.785 6.5 34.209 6.5 56.5C6.5 80.483 26.017 100 50 100s43.5-19.517 43.5-43.5a43.22 43.22 0 0 0-6.72-23.182l4.238-3.431l1.888 2.332a2 2 0 0 0 2.813.297l3.11-2.518a2 2 0 0 0 .294-2.812L89.055 14.75a2 2 0 0 0-2.813-.297l-3.11 2.518a2 2 0 0 0-.294 2.812l1.889 2.332l-4.22 3.414C73.77 18.891 64.883 14.435 55 13.297V8h3a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H42zm8 20c20.2 0 36.5 16.3 36.5 36.5S70.2 93 50 93S13.5 76.7 13.5 56.5S29.8 20 50 20zm.002 7.443L50 56.5l23.234 17.447a29.056 29.056 0 0 0 2.758-30.433a29.056 29.056 0 0 0-25.99-16.07z" color="currentColor"/>`),
		g.Group(children),
	)
}