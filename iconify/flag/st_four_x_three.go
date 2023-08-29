package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#12ad2b" d="M0 0h640v480H0z"/><path fill="#ffce00" d="M0 137.1h640V343H0z"/><path fill="#d21034" d="M0 0v480l240-240"/><g id="flagSt4x30" transform="translate(351.6 240) scale(.34286)"><g id="flagSt4x31"><path id="flagSt4x32" fill="#000" d="M0-200V0h100" transform="rotate(18 0 -200)"/><use width="100%" height="100%" href="#flagSt4x32" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagSt4x31" transform="rotate(72)"/><use width="100%" height="100%" href="#flagSt4x31" transform="rotate(144)"/><use width="100%" height="100%" href="#flagSt4x31" transform="rotate(-144)"/><use width="100%" height="100%" href="#flagSt4x31" transform="rotate(-72)"/></g><use width="100%" height="100%" x="700" href="#flagSt4x30" transform="translate(-523.2)"/>`),
		g.Group(children),
	)
}