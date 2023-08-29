package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="50" cy="50" r="50" fill="#6BC9F2"/><clipPath id="flatUiFlag0"><circle cx="50" cy="50" r="50"/></clipPath><path fill="#fff" fill-rule="evenodd" d="M29 99a2.5 2.5 0 1 1-5 0V17.5a2.5 2.5 0 1 1 5 0V99z" clip-path="url(#flatUiFlag0)" clip-rule="evenodd"/><path fill="#F29C1F" d="M100 73H60V25h40l-7 24z"/><path fill="#E57E25" fill-rule="evenodd" d="M60 73V27l9-8v46l-9 8z" clip-rule="evenodd"/><path fill="#fff" d="m77.017 43.501l-10.791 3.625L65.933 58l-6.962-8.633L48 52.463l6.488-8.962L48 34.539l10.972 3.096L65.933 29l.293 10.875l10.791 3.626z" opacity=".6"/><path fill="#F0C419" d="M29 17h40v48H29z"/><path fill="#fff" d="m69 30.157l-6.028 7.478L52 34.539l6.488 8.962L52 52.463l10.972-3.097L69 56.843z"/>`),
		g.Group(children),
	)
}