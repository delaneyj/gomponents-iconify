package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeverinstallIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#5B29FF" d="M85.661 99.099L0 148.133V49.868z"/><path fill="#2962FF" d="m85.662 99.077l84.676-49.23L256 99.077v97.083l-85.662 49.231v-97.083z"/><path fill="#FFC629" d="M85.661 0L0 49.821l85.661 49.231l84.677-49.231z"/><path fill="#FAFAFA" d="m0 49.821l85.66 49.23l84.678-49.23L85.66 0L0 49.821Zm15.983-.041l69.655 40.032l68.854-40.032L85.637 9.268L15.983 49.78Z"/><path fill="#FFBC00" d="M85.654 89.856V9.34L16 49.849z"/>`),
		g.Group(children),
	)
}