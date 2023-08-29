package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndCallRoundedBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m6.947 16.517l-1.34.38C3.782 17.415 2 15.91 2 13.85c0-1.237.277-2.477 1.083-3.347C4.128 9.376 6 7.908 9 7.292v6.326c0 1.365-.844 2.556-2.053 2.9ZM15 13.618c0 1.365.844 2.556 2.053 2.9l1.34.38C20.218 17.414 22 15.91 22 13.85c0-1.237-.277-2.477-1.083-3.347C19.872 9.376 18 7.908 15 7.292v6.326Z" clip-rule="evenodd"/><path d="M9 13.618s0-1.654 3-1.654s3 1.654 3 1.654V7.292A14.886 14.886 0 0 0 12 7c-1.106 0-2.103.108-3 .292v6.326Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}