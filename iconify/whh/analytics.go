package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analytics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 448V0q91 0 174 35.5T893 131t95.5 143t35.5 174H576zm-128 576q-91 0-174-35.5T131 893T35.5 750T0 576t35.5-174T131 259t143-95.5T448 128v448h448q0 91-35.5 174T765 893t-143 95.5t-174 35.5z"/>`),
		g.Group(children),
	)
}