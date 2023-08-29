package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M272 191.997l48 48-48 47.999-48-48z" fill="currentColor"/><path d="M208.007 256l48.001 48L80.01 480l-48-48.001z" fill="currentColor"/><path d="M256 32h32v64h-32z" fill="currentColor"/><path d="M385.131 149.497l-22.627-22.627 45.254-45.254 22.627 22.628z" fill="currentColor"/><path d="M113.606 104.234l22.627-22.627 45.255 45.254-22.628 22.627z" fill="currentColor"/><path d="M362.505 353.126l22.627-22.627 45.255 45.254-22.628 22.627z" fill="currentColor"/><path d="M416 224h64v32h-64z" fill="currentColor"/>`),
		g.Group(children),
	)
}