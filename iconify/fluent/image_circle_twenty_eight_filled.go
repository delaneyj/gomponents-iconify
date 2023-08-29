package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCircleTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2C7.373 2 2 7.373 2 14a11.96 11.96 0 0 0 3.016 7.956l7.588-7.4a2 2 0 0 1 2.792 0l7.588 7.4A11.955 11.955 0 0 0 26 14c0-6.627-5.373-12-12-12Zm7.921 21.014l-7.572-7.383a.5.5 0 0 0-.698 0l-7.572 7.383A11.954 11.954 0 0 0 14 26c3.035 0 5.808-1.127 7.921-2.986ZM20.5 9.75a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0Z"/>`),
		g.Group(children),
	)
}