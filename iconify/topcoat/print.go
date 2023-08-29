package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M12.5 30.504V33.5h15v-2.996h-15zm0-5V28.5h15v-2.996h-15zm-9 2.996h3v8.002c-.01 2.43.548 2.998 3 2.998h21c2.452 0 3-.452 3-2.998V28.5h3c2.55.04 3.029-.527 3-2.998v-11c0-2.35-.38-3.002-3-3.002h-3V6.502c.029-2.47-.45-3.002-3-3.002h-21c-2.4 0-3.01.572-3 3.002V11.5h-3c-2.58 0-3 .562-3 3.002v11c-.01 2.43.6 3.038 3 2.998zm27 8h-21v-13h21v13zm1.88-18.498c0-.9.72-1.619 1.62-1.619s1.62.719 1.62 1.619s-.72 1.62-1.62 1.62s-1.62-.72-1.62-1.62zM9.5 6.5h21v5.004h-21V6.5z"/>`),
		g.Group(children),
	)
}