package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSymlinkFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M8.5 1H1c-.55 0-1 .45-1 1v12c0 .55.45 1 1 1h10c.55 0 1-.45 1-1V4.5L8.5 1zM11 14H1V2h7l3 3v9zM6 4.5l4 3l-4 3v-2c-.98-.02-1.84.22-2.55.7c-.71.48-1.19 1.25-1.45 2.3c.02-1.64.39-2.88 1.13-3.73c.73-.84 1.69-1.27 2.88-1.27v-2H6z" fill="currentColor"/>`),
		g.Group(children),
	)
}