package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 15c-2.229 0-6-.421-6-2c0-.04.004-.079.014-.117l.047-.2a.504.504 0 0 1 .036-.1l5.453-11.3a.52.52 0 0 1 .9 0l5.453 11.3a.5.5 0 0 1 .036.1l.047.2c.01.038.014.077.014.117c0 1.579-3.771 2-6 2Z"/>`),
		g.Group(children),
	)
}