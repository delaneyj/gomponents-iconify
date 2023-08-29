package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KitchenKnife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m14.665 10.393l1.387 1.492l.003-.002l13.454 14.404l.296.316a.744.744 0 0 1 .14.215a.733.733 0 0 1-.18.825l-.34.32a7.763 7.763 0 0 1-10.98-.38l-8.41-9.01c-.3-.32-.28-.82.04-1.11l2.362-2.204l-1.382-1.496l-3.3-3.54a5.154 5.154 0 0 0-2.94-1.57a3.358 3.358 0 0 1-1.68-5.83c1.38-1.21 3.48-1.06 4.73.28l6.8 7.29Zm14.028 16.485L16.006 13.296L11.4 17.594l7.467 8a7.786 7.786 0 0 0 9.826 1.284ZM8.035 5.903a.91.91 0 1 0-1.82 0a.91.91 0 0 0 1.82 0Zm3.64 3.89a.91.91 0 1 0-1.82 0a.91.91 0 0 0 1.82 0Z"/>`),
		g.Group(children),
	)
}