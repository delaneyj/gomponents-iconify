package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Split(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.786 20.698c-1.792-.237-2.912-1.33-4.358-2.886c-.697-.75-1.428-1.577-2.324-2.32c1.396-1.164 2.41-2.518 3.483-3.502c1.01-.92 1.9-1.52 3.2-1.688v2.574l7.555-4.363l-7.556-4.363v2.652c-3.34.266-5.45 2.378-6.934 4.013c-.82.896-1.537 1.692-2.212 2.192c-.685.5-1.227.73-2.013.742H2.812v3.5h7.814c.786.01 1.33.24 2.017.742c1.02.743 2.095 2.18 3.552 3.568c1.312 1.258 3.162 2.46 5.592 2.65v2.663l7.556-4.36l-7.556-4.36v2.545z"/>`),
		g.Group(children),
	)
}