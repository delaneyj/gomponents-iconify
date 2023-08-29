package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayConsoleAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.183 9.873l-8.94-5.162A1.558 1.558 0 0 0 7.905 6.06v35.88a1.558 1.558 0 0 0 2.336 1.349l31.073-17.94a1.558 1.558 0 0 0 0-2.698l-2.74-1.583"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.913 34.241H14.366a1.558 1.558 0 0 1-1.557-1.557V9.874h25.765V26.93M12.809 15.245h25.765"/><circle cx="18.529" cy="26.817" r="1.558" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.788" cy="21.559" r="1.558" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="27.526" cy="25.298" r="1.558" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.14" cy="19.685" r="1.558" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.038 20.786l-3.41 3.411m-3.739-1.536l1.536 1.536m-6.794 1.519l3.055-3.055"/>`),
		g.Group(children),
	)
}