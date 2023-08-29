package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M40.588 59.659L27.579 82.03h46.912L87.5 59.659zm45.364-1.28L62.496 17.753l-25.879-.081L60.073 58.3zm-49.996-39.13L12.5 59.876l12.87 22.452l23.456-40.627z"/>`),
		g.Group(children),
	)
}