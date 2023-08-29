package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NaturalFeature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M12 1h9v7h-9zm18.723 28.822c-.873 0-1.723-.891-1.723-1.77V10h14.187l5.829 24.878L49 36.354V47.41c0 .85-.707 1.59-1.549 1.59H34.03c-.833 0-1.03-.74-1.03-1.59V30h-2m-11.719-.158c.866 0 1.719-.859 1.719-1.744V10H6.822L.986 34.892L1 36.4v11.053C1 48.307 1.707 49 2.548 49h13.405c.854 0 1.047-.693 1.047-1.547V30h2m4-20h4v17h-4zm6-9h10v7H29z"/>`),
		g.Group(children),
	)
}