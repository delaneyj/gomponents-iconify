package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosFitness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M352 280H160c-4.4 0-8-3.6-8-8v-32c0-4.4 3.6-8 8-8h192c4.4 0 8 3.6 8 8v32c0 4.4-3.6 8-8 8z" fill="currentColor"/><path d="M378 362V150c0-12.1 9.9-22 22-22s22 9.9 22 22v212c0 12.1-9.9 22-22 22s-22-9.9-22-22z" fill="currentColor"/><path d="M460 192h-12c-4.4 0-8 3.6-8 8v112c0 4.4 3.6 8 8 8h12c11 0 20-9 20-20v-88c0-11-9-20-20-20z" fill="currentColor"/><path d="M134 362V150c0-12.1-9.9-22-22-22s-22 9.9-22 22v212c0 12.1 9.9 22 22 22s22-9.9 22-22z" fill="currentColor"/><path d="M64 192H52c-11 0-20 9-20 20v88c0 11 9 20 20 20h12c4.4 0 8-3.6 8-8V200c0-4.4-3.6-8-8-8z" fill="currentColor"/>`),
		g.Group(children),
	)
}