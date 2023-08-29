package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coursera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.968 24.945h2.412m.62 1.861l-1.791-5.512l-1.861 5.512m3.032-1.861h-2.343m-10.051-3.651v5.512m0-2.756h1.792m-1.792-2.756h2.756m-2.756 5.512h2.756m-8.039-.62a1.643 1.643 0 0 0 1.378.62h.826a1.378 1.378 0 0 0 0-2.756h-.895a1.378 1.378 0 0 1 0-2.756h.826a1.48 1.48 0 0 1 1.378.62m8.061 3.025l1.797 1.867m-3.637-.081v-5.53h1.798a1.867 1.867 0 0 1 0 3.733h-1.798m-12.875.011l1.798 1.867m-3.638-.081v-5.53h1.798a1.867 1.867 0 0 1 0 3.733h-1.798m-4.864-3.634v3.652a1.792 1.792 0 1 0 3.583 0v-3.652m-8.502 3.672a1.806 1.806 0 1 0 3.611 0v-1.84a1.832 1.832 0 0 0-1.84-1.84a1.775 1.775 0 0 0-1.771 1.84Zm-1.259 0a1.84 1.84 0 0 1-3.68 0v-1.84a1.832 1.832 0 0 1 1.84-1.84a1.775 1.775 0 0 1 1.772 1.84"/>`),
		g.Group(children),
	)
}