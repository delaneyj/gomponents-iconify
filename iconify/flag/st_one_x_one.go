package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#12ad2b" d="M0 0h512v512H0z"/><path fill="#ffce00" d="M0 146.3h512v219.4H0z"/><path fill="#d21034" d="M0 0v512l192-256"/><g id="flagSt1x10" transform="translate(276.9 261.5) scale(.33167)"><g id="flagSt1x11"><path id="flagSt1x12" fill="#000" d="M0-200V0h100" transform="rotate(18 0 -200)"/><use width="100%" height="100%" href="#flagSt1x12" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagSt1x11" transform="rotate(72)"/><use width="100%" height="100%" href="#flagSt1x11" transform="rotate(144)"/><use width="100%" height="100%" href="#flagSt1x11" transform="rotate(-144)"/><use width="100%" height="100%" href="#flagSt1x11" transform="rotate(-72)"/></g><use width="100%" height="100%" x="700" href="#flagSt1x10" transform="translate(-550.9)"/>`),
		g.Group(children),
	)
}