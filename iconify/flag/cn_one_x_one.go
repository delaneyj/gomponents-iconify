package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CnOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><path id="flagCn1x10" fill="#ff0" d="M1-.3L-.7.8L0-1L.6.8L-1-.3z"/></defs><path fill="#ee1c25" d="M0 0h512v512H0z"/><use width="30" height="20" href="#flagCn1x10" transform="matrix(76.8 0 0 76.8 128 128)"/><use width="30" height="20" href="#flagCn1x10" transform="rotate(-121 142.6 -47) scale(25.5827)"/><use width="30" height="20" href="#flagCn1x10" transform="rotate(-98.1 198 -82) scale(25.6)"/><use width="30" height="20" href="#flagCn1x10" transform="rotate(-74 272.4 -114) scale(25.6137)"/><use width="30" height="20" href="#flagCn1x10" transform="matrix(16 -19.968 19.968 16 256 230.4)"/>`),
		g.Group(children),
	)
}