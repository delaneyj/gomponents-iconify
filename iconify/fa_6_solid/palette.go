package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 256v2.7c-.4 36.5-33.6 61.3-70.1 61.3H344c-26.5 0-48 21.5-48 48c0 3.4.4 6.7 1 9.9c2.1 10.2 6.5 20 10.8 29.9c6.1 13.8 12.1 27.5 12.1 42c0 31.8-21.6 60.7-53.4 62c-3.5.1-7 .2-10.6.2C114.6 512 0 397.4 0 256S114.6 0 256 0s256 114.6 256 256zm-384 32a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm0-96a32 32 0 1 0 0-64a32 32 0 1 0 0 64zm160-96a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm96 96a32 32 0 1 0 0-64a32 32 0 1 0 0 64z"/>`),
		g.Group(children),
	)
}