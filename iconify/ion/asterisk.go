package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M480 224l-186.828 7.487L401.688 64l-59.247-32L256 208 169.824 32l-59.496 32 108.5 167.487L32 224v64l185.537-10.066L113.65 448l55.969 32L256 304l86.381 176 55.949-32-103.867-170.066L480 288z" fill="currentColor"/>`),
		g.Group(children),
	)
}