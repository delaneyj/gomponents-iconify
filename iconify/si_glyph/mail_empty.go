package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.263 5.806L8.996.063L1.689 5.806c-.358 0-.648.3-.648.667v8.857c0 .367.29.667.648.667h14.573c.358 0 .648-.3.648-.667V6.473a.656.656 0 0 0-.647-.667zm-1.279 9.225l-2.826-3.945l-3.162 2.622l-3.134-2.607l-2.878 3.931h-.995L5 10.207L2.045 7.156l.014-.669l6.938-5.228L15.95 6.52v.636L13.014 10.2l3.002 4.817l-1.032.014z"/>`),
		g.Group(children),
	)
}