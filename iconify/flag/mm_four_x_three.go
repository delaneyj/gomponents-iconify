package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MmFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fecb00" d="M0 0h640v480H0z"/><path fill="#34b233" d="M0 160h640v320H0z"/><path fill="#ea2839" d="M0 320h640v160H0z"/><g transform="translate(320 256.9) scale(176.87999)"><path id="flagMm4x30" fill="#fff" d="m0-1l.3 1h-.6z"/><use width="100%" height="100%" href="#flagMm4x30" transform="rotate(-144)"/><use width="100%" height="100%" href="#flagMm4x30" transform="rotate(-72)"/><use width="100%" height="100%" href="#flagMm4x30" transform="rotate(72)"/><use width="100%" height="100%" href="#flagMm4x30" transform="rotate(144)"/></g>`),
		g.Group(children),
	)
}