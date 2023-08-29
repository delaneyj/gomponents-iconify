package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransgenderNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsTransgenderNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm14 21c0 5.185 3.947 9.449 9 9.95v5.98h-2a1 1 0 1 0 0 2h2V41a1 1 0 1 0 2 0v-2.07h2a1 1 0 0 0 0-2h-2v-5.98c5.053-.501 9-4.765 9-9.95a9.959 9.959 0 0 0-1.073-4.511l2.104-2.091l1.43 1.43a1 1 0 0 0 1.414-1.415l-1.425-1.425L39 10.454V12a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1h-4a1 1 0 0 0 0 2h1.626l-2.59 2.574l-1.404-1.404a1 1 0 1 0-1.414 1.415l1.4 1.398l-1.797 1.785A9.981 9.981 0 0 0 24 11a9.981 9.981 0 0 0-7.82 3.766l-.013-.013L10.414 9H12a1 1 0 1 0 0-2H8a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-1.586l5.753 5.753a.997.997 0 0 0 .365.233A9.957 9.957 0 0 0 14 21Zm10-8a8 8 0 1 0 0 16a8 8 0 0 0 0-16Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsTransgenderNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}