package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdCodeDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M234.6 160v125.7l-44.7-43.6L160 272l96 96 96-96-29.9-31-44.7 44.7V160h-42.8z" fill="currentColor"/><path d="M190.4 354.1L91.9 256l98.4-98.1-30-29.9L32 256l128.4 128 30-29.9zm131.2 0L420 256l-98.4-98.1 30-29.9L480 256 351.6 384l-30-29.9z" fill="currentColor"/>`),
		g.Group(children),
	)
}