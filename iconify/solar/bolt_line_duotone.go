package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoltLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M13.926 9.705c-.379-.371-.379-.963-.379-2.148v-.31c0-3.285 0-4.927-.923-5.21c-.923-.283-1.913 1.056-3.892 3.734L5.67 9.914c-1.285 1.739-1.928 2.608-1.574 3.291l.018.034c.375.673 1.485.673 3.704.673c1.233 0 1.85 0 2.236.363"/><path d="m13.926 9.706l.02.019c.387.364 1.003.364 2.236.364c2.22 0 3.329 0 3.703.672l.019.034c.354.684-.289 1.553-1.574 3.29l-3.062 4.144c-1.98 2.678-2.969 4.017-3.892 3.734c-.924-.283-.924-1.925-.923-5.21v-.31c0-1.184 0-1.777-.379-2.148l-.02-.02" opacity=".5"/></g>`),
		g.Group(children),
	)
}