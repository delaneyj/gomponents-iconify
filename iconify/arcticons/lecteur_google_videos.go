package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LecteurGoogleVideos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 35.584c2.051-1.766 6.804-5.13 9.621-5.13c2.692 0 18.837 7.06 25.379 12.046"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.628 32.142c2.051-1.766 8.235-4.915 11.053-4.915c2.576 0 5.857 2.468 10.819 5.553"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.667 27.627c.254 1.228-1.176 4.256-.672 6.738a27.058 27.058 0 0 0 1.822 4.518m-18.715-8.277c-3.63 2.65 1.417 8.548 5.202 11.894"/><circle cx="15.374" cy="15.818" r="3.449" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}