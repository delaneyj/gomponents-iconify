package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportGeneralSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M13.727 2.98c.335.403.605.842.808 1.303a2 2 0 0 1-2.551.44l1.743-1.743z" fill="currentColor"/><path d="M13.02 2.273l-1.743 1.743a2 2 0 0 1 .44-2.55c.461.202.9.472 1.302.807z" fill="currentColor"/><path d="M14.12 9.485a5.506 5.506 0 0 0 .757-4.144a3.002 3.002 0 0 1-3.616.105l-.59.59l3.45 3.449z" fill="currentColor"/><path d="M10.554 4.739a3.002 3.002 0 0 1 .105-3.616a5.506 5.506 0 0 0-4.144.756l3.45 3.45l.59-.59z" fill="currentColor"/><path d="M5.723 2.502a5.484 5.484 0 0 0-1.297 1.87c1.334.39 2.948 1.06 4.273 2.222l.558-.558l-3.534-3.534z" fill="currentColor"/><path d="M9.409 7.298l.555-.555l3.534 3.534a5.51 5.51 0 0 1-1.848 1.287c-.391-1.336-1.066-2.946-2.241-4.266z" fill="currentColor"/><path d="M11.017 13.374c-.001 1.027-.925 1.811-1.995 1.583c-1.479-.314-3.726-1.025-5.318-2.616c-1.587-1.587-2.31-3.836-2.636-5.32a1.783 1.783 0 0 1-.043-.383c0-1.024.92-1.812 1.993-1.587c1.496.312 3.769 1.02 5.362 2.613c1.597 1.596 2.295 3.863 2.6 5.35c.025.123.037.243.037.36zM5.603 8.897a.5.5 0 0 0-.707.707l1.5 1.5a.5.5 0 1 0 .707-.707l-1.5-1.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}