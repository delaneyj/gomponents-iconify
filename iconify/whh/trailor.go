package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trailor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 768h-86q-28-58-82-93t-120-35t-120 35t-82 93H32q-13 0-22.5-9.5T0 736t9.5-22.5T32 704h96V256q0-106 75-181T384 0h384q106 0 181 75t75 181v448q0 26-18.5 45T960 768zM448 256q0-27-19-45.5T384 192h-64q-27 0-45.5 18.5T256 256v64q0 26 18.5 45t45.5 19h64q27 0 45.5-19t18.5-45v-64zm448 0q0-27-18.5-45.5T832 192H576q-26 0-45 18.5T512 256v64q0 26 19 45t45 19h256q27 0 45.5-19t18.5-45v-64zM672 704q66 0 113 47t47 113t-47 113t-113 47t-113-47t-47-113t47-113t113-47zm0 192q13 0 22.5-9.5T704 864t-9.5-22.5T672 832t-22.5 9.5T640 864t9.5 22.5T672 896z"/>`),
		g.Group(children),
	)
}