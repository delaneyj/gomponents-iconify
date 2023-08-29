package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerNoneTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12.92 3.316c.774-.69 1.983-.187 2.074.812L15 4.25v15.496c0 1.037-1.178 1.606-1.986 1.01l-.095-.076l-4.491-3.994a.75.75 0 0 0-.39-.182l-.108-.008H4.25a2.25 2.25 0 0 1-2.245-2.095L2 14.246V9.75a2.25 2.25 0 0 1 2.096-2.245l.154-.005h3.68a.75.75 0 0 0 .411-.123l.087-.067l4.491-3.993zm4.36 5.904L19 10.94l1.72-1.72a.75.75 0 1 1 1.06 1.06L20.06 12l1.72 1.72a.75.75 0 1 1-1.06 1.06L19 13.06l-1.72 1.72a.75.75 0 1 1-1.06-1.06L17.94 12l-1.72-1.72a.75.75 0 1 1 1.06-1.06z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}