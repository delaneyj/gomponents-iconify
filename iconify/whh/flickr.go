package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800 568q-93 0-158.5-65.5T576 344t65.5-158.5T800 120t158.5 65.5T1024 344t-65.5 158.5T800 568zm-576 0q-93 0-158.5-65.5T0 344t65.5-158.5T224 120t158.5 65.5T448 344t-65.5 158.5T224 568z"/>`),
		g.Group(children),
	)
}