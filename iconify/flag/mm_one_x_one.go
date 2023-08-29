package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MmOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fecb00" d="M0 0h512v512H0z"/><path fill="#34b233" d="M0 170.7h512V512H0z"/><path fill="#ea2839" d="M0 341.3h512V512H0z"/><path id="flagMm1x10" fill="#fff" stroke-width="188.7" d="M312.6 274H199.4L256 85.3Z"/><use width="100%" height="100%" href="#flagMm1x10" transform="rotate(-144 256 274)"/><use width="100%" height="100%" href="#flagMm1x10" transform="rotate(-72 256 274)"/><use width="100%" height="100%" href="#flagMm1x10" transform="rotate(72 256 274)"/><use width="100%" height="100%" href="#flagMm1x10" transform="rotate(144 256 274)"/>`),
		g.Group(children),
	)
}