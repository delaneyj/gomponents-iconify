package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithTongue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="20.5" cy="24.5" r="5" fill="currentColor"/><circle cx="43.5" cy="24.5" r="5" fill="currentColor"/><path fill="currentColor" d="M32 2C15.431 2 2 15.432 2 32s13.431 30 30 30c16.568 0 30-13.432 30-30S48.568 2 32 2m.002 58c-5.521-.004-10-4.398-10-13.156v-3.995c0-.531.113-.849.822-.849H30.5l1.502 13.789L33.5 42h7.68c.708 0 .822.317.822.849v3.995c0 8.758-4.479 13.152-10 13.156m8.486-1.856c1.647-1.895 2.881-4.573 3.328-8.203c3.394-2.907 5.187-6.981 5.187-10.94c0-.493-.392-1-1.125-1H16.125C15.392 38 15 38.507 15 39c0 3.957 1.792 8.029 5.183 10.937c.446 3.638 1.69 6.313 3.342 8.211C12.498 54.564 4.5 44.205 4.5 32C4.5 16.836 16.836 4.5 32 4.5c15.163 0 27.5 12.336 27.5 27.5c0 12.201-7.992 22.557-19.012 26.144"/>`),
		g.Group(children),
	)
}