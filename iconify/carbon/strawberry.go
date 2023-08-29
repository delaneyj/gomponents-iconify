package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strawberry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 10h2v2H9zm0 4h2v2H9zm4-2h2v2h-2z"/><path fill="currentColor" d="M9 28a5.169 5.169 0 0 1-.744-.054A5.093 5.093 0 0 1 4 22.866V11.199a7.162 7.162 0 0 1 6.31-7.165a6.96 6.96 0 0 1 4.464 1.07l7.97 4.981a6.98 6.98 0 0 1-.69 12.19l-10.88 5.229A4.976 4.976 0 0 1 9 28zm2.002-22q-.248 0-.5.024A5.146 5.146 0 0 0 6 11.198v11.669a3.084 3.084 0 0 0 2.543 3.1a3.027 3.027 0 0 0 1.763-.265l10.883-5.23a4.98 4.98 0 0 0 .488-8.696l-7.972-4.982A4.969 4.969 0 0 0 11.002 6zM26 7h-4a4.005 4.005 0 0 1-4-4V2h2v1a2.002 2.002 0 0 0 2 2h4z"/>`),
		g.Group(children),
	)
}