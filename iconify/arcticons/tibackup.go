package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tibackup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20 4a1 1 0 0 0-1 .84l-.75 5.3a14.8 14.8 0 0 0-3.37 2l-5-2a1 1 0 0 0-1.23.43l-4 6.93a1 1 0 0 0 .24 1.28l4.21 3.31A18.26 18.26 0 0 0 9 24a15.7 15.7 0 0 0 .13 2l-4.21 3.26a1 1 0 0 0-.24 1.28l4 6.93a1 1 0 0 0 1.23.43l5-2a15.11 15.11 0 0 0 3.37 2l.75 5.3a1 1 0 0 0 .97.8h8a1 1 0 0 0 1-.84l.76-5.3a14.8 14.8 0 0 0 3.37-2l5 2a1 1 0 0 0 1.22-.43l4-6.93a1 1 0 0 0-.24-1.28L38.85 26a15.7 15.7 0 0 0 .15-2a15.7 15.7 0 0 0-.14-1.95l4.23-3.31a1 1 0 0 0 .25-1.28l-4-6.93a1 1 0 0 0-1.22-.43l-5 2a15.18 15.18 0 0 0-3.38-2L29 4.84A1 1 0 0 0 28 4Zm7 25.49h5.49M27 18.5h5.49m-2.74 0v10.99m-12.84 0h5.49M16.91 18.5h5.49m-2.74 0v10.99M15.51 18.5h8.3m-4.15 11v-11m-4.15 0v2.73m8.3-2.73v2.73"/>`),
		g.Group(children),
	)
}