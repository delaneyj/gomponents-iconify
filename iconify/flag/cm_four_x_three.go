package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CmFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#007a5e" d="M0 0h213.3v480H0z"/><path fill="#ce1126" d="M213.3 0h213.4v480H213.3z"/><path fill="#fcd116" d="M426.7 0H640v480H426.7z"/><g fill="#fcd116" transform="translate(320 240) scale(7.1111)"><g id="flagCm4x30"><path id="flagCm4x31" d="M0-8L-2.5-.4L1.3.9z"/><use width="100%" height="100%" href="#flagCm4x31" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagCm4x30" transform="rotate(72)"/><use width="100%" height="100%" href="#flagCm4x30" transform="rotate(144)"/><use width="100%" height="100%" href="#flagCm4x30" transform="rotate(-144)"/><use width="100%" height="100%" href="#flagCm4x30" transform="rotate(-72)"/></g>`),
		g.Group(children),
	)
}