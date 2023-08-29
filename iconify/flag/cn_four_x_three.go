package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CnFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><path id="flagCn4x30" fill="#ff0" d="M-.6.8L0-1L.6.8L-1-.3h2z"/></defs><path fill="#ee1c25" d="M0 0h640v480H0z"/><use width="30" height="20" href="#flagCn4x30" transform="matrix(71.9991 0 0 72 120 120)"/><use width="30" height="20" href="#flagCn4x30" transform="matrix(-12.33562 -20.5871 20.58684 -12.33577 240.3 48)"/><use width="30" height="20" href="#flagCn4x30" transform="matrix(-3.38573 -23.75998 23.75968 -3.38578 288 95.8)"/><use width="30" height="20" href="#flagCn4x30" transform="matrix(6.5991 -23.0749 23.0746 6.59919 288 168)"/><use width="30" height="20" href="#flagCn4x30" transform="matrix(14.9991 -18.73557 18.73533 14.99929 240 216)"/>`),
		g.Group(children),
	)
}