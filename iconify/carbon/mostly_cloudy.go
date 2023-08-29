package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MostlyCloudy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.743 18.692a6.05 6.05 0 0 0 1.057-1.086a5.998 5.998 0 1 0-10.733-4.445A7.557 7.557 0 0 0 6.35 18.25A5.993 5.993 0 0 0 8 30.005h11a5.985 5.985 0 0 0 2.743-11.313ZM18 10.005a4.004 4.004 0 0 1 4 4a3.961 3.961 0 0 1-.8 2.4a4.012 4.012 0 0 1-.94.891a7.542 7.542 0 0 0-6.134-4.24A3.998 3.998 0 0 1 18 10.006Zm1 18H8a3.993 3.993 0 0 1-.673-7.93l.663-.112l.146-.656a5.496 5.496 0 0 1 10.729 0l.146.656l.662.112a3.993 3.993 0 0 1-.673 7.93Zm7-15h4v2h-4zM22.95 7.64l2.828-2.827l1.415 1.414l-2.829 2.828zM17 2.005h2v4h-2zM8.808 6.227l1.414-1.414l2.829 2.828l-1.415 1.414z"/>`),
		g.Group(children),
	)
}