package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kdeconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.22 4.69l-6.4.62V32l6.33-1V19.66l8.51 12.63l6.67-2.14L33.62 18L42.4 6.55L35.59 5l-8.44 11.42ZM12.71 11.2a.33.33 0 0 0-.2.09L10 13.84a.35.35 0 0 0-.05.4l2.94 4.91A13 13 0 0 0 11.66 22l-5.4 1.13a.34.34 0 0 0-.26.33v3.6a.33.33 0 0 0 .25.32l5.24 1.29A12.92 12.92 0 0 0 12.77 32l-3 4.69a.33.33 0 0 0 0 .42l2.51 2.54a.32.32 0 0 0 .4 0l4.75-2.93a12.32 12.32 0 0 0 3 1.27l1.11 5.41a.32.32 0 0 0 .32.26h3.55a.32.32 0 0 0 .31-.25L27.05 38a12.47 12.47 0 0 0 3.09-1.32l4.68 3.12a.33.33 0 0 0 .41 0l2.51-2.54a.33.33 0 0 0 0-.41L36.08 34l-.55.18c-.08 0 .42-1.39.13-1.86l-2.19-3.24c-1 4.09-5.77 6.94-9.72 6.94a10.25 10.25 0 0 1-10.18-10.34c0-4.2 3.27-9.2 7.25-9.86v-2.66a7.56 7.56 0 0 0-3.1 1.28h0l-4.8-3.18a.27.27 0 0 0-.21-.05Z"/>`),
		g.Group(children),
	)
}