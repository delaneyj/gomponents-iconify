package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailRewindTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.998 7.999V6.8l2.36 1.966a1 1 0 0 0 1.64-.768V3.002a1 1 0 0 0-1.64-.768L15.998 4.2V3.002a1 1 0 0 0-1.64-.768L11.36 4.732a1 1 0 0 0 0 1.537l2.998 2.498a1 1 0 0 0 1.64-.768Zm-2.28 1.536c.089.074.181.14.276.195l-3.74 2.201a.5.5 0 0 1-.426.038l-.082-.038L3 7.963V14.5A1.5 1.5 0 0 0 4.5 16h11a1.5 1.5 0 0 0 1.5-1.5V8.937l.717.598c.092.076.186.142.283.2V14.5a2.5 2.5 0 0 1-2.5 2.5h-11A2.5 2.5 0 0 1 2 14.5v-8A2.5 2.5 0 0 1 4.5 4h6.177a1.985 1.985 0 0 0-.614 1H4.5A1.5 1.5 0 0 0 3 6.5v.302l7 4.118l3.153-1.855l.565.47Z"/>`),
		g.Group(children),
	)
}