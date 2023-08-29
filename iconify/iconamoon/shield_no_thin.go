package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldNoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.197-.46a.5.5 0 0 0-.394 0L12 3Zm0 18l-.248.434a.5.5 0 0 0 .496 0L12 21Zm6.394-15.26l-.197.46l.197-.46ZM8.024 18.727l-.249.435l.248-.434ZM11.802 2.54L5.409 5.28l.394.92l6.394-2.74l-.394-.92ZM4.5 6.66v6.858h1V6.66h-1Zm3.275 12.502l3.977 2.272l.496-.868l-3.977-2.273l-.496.869Zm4.473 2.272l3.977-2.272l-.496-.869l-3.977 2.273l.496.868Zm7.252-7.916V6.66h-1v6.86h1Zm-.91-8.237l-6.393-2.74l-.394.919l6.394 2.74l.394-.92Zm.91 1.378a1.5 1.5 0 0 0-.91-1.378l-.393.919a.5.5 0 0 1 .303.46h1Zm-3.275 12.503a6.5 6.5 0 0 0 3.275-5.644h-1a5.5 5.5 0 0 1-2.771 4.775l.496.869ZM4.5 13.518a6.5 6.5 0 0 0 3.275 5.644l.496-.869A5.5 5.5 0 0 1 5.5 13.518h-1Zm.91-8.237a1.5 1.5 0 0 0-.91 1.378h1c0-.2.12-.38.303-.46l-.394-.918Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10 10l4 4m-4 0l4-4"/></g>`),
		g.Group(children),
	)
}