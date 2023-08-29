package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteForeverTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M19.5 16a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zM14 2a4 4 0 0 1 3.995 3.8L18 6h5a1 1 0 0 1 .117 1.993L23 8h-.849l-.594 7.332a6.5 6.5 0 0 0-6.746 10.669L10.766 26a3.75 3.75 0 0 1-3.738-3.447L5.848 8H5a1 1 0 0 1-.993-.883L4 7a1 1 0 0 1 .883-.993L5 6h5a4 4 0 0 1 4-4zm3.73 17.024l-.068-.058a.5.5 0 0 0-.569 0l-.07.058l-.057.07a.5.5 0 0 0 0 .568l.058.07l1.77 1.769l-1.768 1.766l-.057.07a.5.5 0 0 0 0 .568l.057.07l.07.057a.5.5 0 0 0 .568 0l.07-.057l1.766-1.767l1.77 1.769l.069.058a.5.5 0 0 0 .568 0l.07-.058l.057-.07a.5.5 0 0 0 0-.568l-.057-.07l-1.77-1.768l1.772-1.77l.058-.069a.5.5 0 0 0 0-.569l-.058-.069l-.069-.058a.5.5 0 0 0-.569 0l-.069.058l-1.772 1.77l-1.77-1.77l-.068-.058l.069.058zM14 4a2 2 0 0 0-1.995 1.85L12 6h4l-.005-.15A2 2 0 0 0 14 4z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}