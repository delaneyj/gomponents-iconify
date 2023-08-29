package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GwOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#ce1126" d="M0 0h160v512H0z"/><path fill="#fcd116" d="M160 0h352v256H160z"/><path fill="#009e49" d="M160 256h352v256H160z"/><g transform="translate(-46.2 72.8) scale(.7886)"><g id="flagGw1x10" transform="matrix(80 0 0 80 160 240)"><path id="flagGw1x11" fill="#000" d="M0-1v1h.5" transform="rotate(18 0 -1)"/><use width="100%" height="100%" href="#flagGw1x11" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagGw1x10" transform="rotate(72 160 240)"/><use width="100%" height="100%" href="#flagGw1x10" transform="rotate(144 160 240)"/><use width="100%" height="100%" href="#flagGw1x10" transform="rotate(-144 160 240)"/><use width="100%" height="100%" href="#flagGw1x10" transform="rotate(-72 160 240)"/></g>`),
		g.Group(children),
	)
}