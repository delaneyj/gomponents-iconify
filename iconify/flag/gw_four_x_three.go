package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GwFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#ce1126" d="M0 0h220v480H0z"/><path fill="#fcd116" d="M220 0h420v240H220z"/><path fill="#009e49" d="M220 240h420v240H220z"/><g id="flagGw4x30" transform="matrix(80 0 0 80 110 240)"><path id="flagGw4x31" fill="#000" d="M0-1v1h.5" transform="rotate(18 0 -1)"/><use width="100%" height="100%" href="#flagGw4x31" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagGw4x30" transform="rotate(72 110 240)"/><use width="100%" height="100%" href="#flagGw4x30" transform="rotate(144 110 240)"/><use width="100%" height="100%" href="#flagGw4x30" transform="rotate(-144 110 240)"/><use width="100%" height="100%" href="#flagGw4x30" transform="rotate(-72 110 240)"/>`),
		g.Group(children),
	)
}