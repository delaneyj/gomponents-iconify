package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M351.9 256L460 193.6l-48-83.2l-108 62.4V48h-96v124.8l-108-62.4l-48 83.2L160.1 256L52 318.4l48 83.2l108-62.4V464h96V339.2l108 62.4l48-83.2L351.9 256z"/>`),
		g.Group(children),
	)
}